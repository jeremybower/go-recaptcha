// Package recaptcha handles reCaptcha (http://www.google.com/recaptcha) form submissions
//
// This package is designed to be called from within an HTTP server or web framework
// which offers reCaptcha form inputs and requires them to be evaluated for correctness
//
// Edit the recaptchaPrivateKey constant before building and using
package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Client is a client for the Recaptcha service.
type Client struct {
	privateKey string
}

// Response is the response from the Recaptcha service.
type Response struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// ServiceURL is the URL of the Recaptcha service.
const ServiceURL = "https://www.google.com/recaptcha/api/siteverify"

func NewClient(privateKey string) *Client {
	return &Client{
		privateKey: privateKey,
	}
}

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func (client *Client) check(remoteip, response string) (*Response, error) {
	resp, err := http.PostForm(
		ServiceURL,
		url.Values{
			"secret":   {client.privateKey},
			"remoteip": {remoteip},
			"response": {response},
		})
	if err != nil {
		log.Printf("Post error: %s\n", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func (client *Client) Confirm(remoteip, response string) (bool, error) {
	r, err := client.check(remoteip, response)
	if err != nil {
		return false, err
	}

	return r.Success, nil
}
