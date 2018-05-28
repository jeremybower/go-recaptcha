// Package recaptcha handles reCaptcha (http://www.google.com/recaptcha) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers reCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the recaptchaPrivateKey constant before building and using
package recaptcha

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Client is a client for the Recaptcha service.
type Client interface {
	Confirm(remoteip, response string) (bool, error)
}

// Options are the configuration options for the client.
type Options struct {
	BaseURL    string
	HTTPClient *http.Client
	PrivateKey string
	ReaderFunc func(io.Reader) io.Reader
}

// NewOptions will create new options with default values.
func NewOptions(privateKey string) *Options {
	return &Options{
		BaseURL:    "https://www.google.com/recaptcha/api/siteverify",
		HTTPClient: &http.Client{},
		PrivateKey: privateKey,
		ReaderFunc: func(r io.Reader) io.Reader {
			return r
		},
	}
}

// NewClient will create a new client with default options.
func NewClient(privateKey string) Client {
	opts := NewOptions(privateKey)
	return NewClientWithOptions(opts)
}

// NewClientWithOptions will create a new client with the given options.
func NewClientWithOptions(opts *Options) Client {
	return &clientImpl{
		opts: opts,
	}
}

type clientImpl struct {
	opts *Options
}

type serverResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func (c *clientImpl) Confirm(remoteip, response string) (bool, error) {
	r, err := c.check(remoteip, response)
	if err != nil {
		return false, err
	}

	return r.Success, nil
}

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func (c *clientImpl) check(remoteip, response string) (*serverResponse, error) {
	resp, err := c.opts.HTTPClient.PostForm(
		c.opts.BaseURL,
		url.Values{
			"secret":   {c.opts.PrivateKey},
			"remoteip": {remoteip},
			"response": {response},
		})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected response. Expected 200 but found %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(c.opts.ReaderFunc(resp.Body))
	if err != nil {
		return nil, err
	}

	var r serverResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
