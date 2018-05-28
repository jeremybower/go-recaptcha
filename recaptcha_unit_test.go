// +build unit

package recaptcha

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupMockServer() (*http.ServeMux, *httptest.Server, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	return mux, server, func() {
		server.Close()
	}
}

func TestTimeoutErrorUsingMockServer(t *testing.T) {
	mux, server, shutdown := setupMockServer()
	defer shutdown()

	ch := make(chan int)
	defer func() { ch <- 0 }()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = <-ch
		w.WriteHeader(http.StatusOK)
	})

	opts := NewOptions("private-key")
	opts.BaseURL = server.URL
	opts.HTTPClient = &http.Client{
		Timeout: 10 * time.Millisecond,
	}
	client := NewClientWithOptions(opts)

	ok, err := client.Confirm("126.0.0.1", "any-recaptcha-response")

	if err == nil {
		t.Error("Expected timeout error")
	}

	if ok {
		t.Error("Expected ok to be false")
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestReadErrorUsingMockServer(t *testing.T) {
	mux, server, shutdown := setupMockServer()
	defer shutdown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	opts := NewOptions("private-key")
	opts.BaseURL = server.URL
	opts.ReaderFunc = func(io.Reader) io.Reader {
		return errReader(0)
	}
	client := NewClientWithOptions(opts)

	ok, err := client.Confirm("126.0.0.1", "any-recaptcha-response")
	expectedError := "test error"
	if err.Error() != expectedError {
		t.Errorf("Unexpected error: %s", err)
	}

	if ok {
		t.Error("Expected ok to be false")
	}
}

func TestUnexpectedStatusCodeUsingMockServer(t *testing.T) {
	mux, server, shutdown := setupMockServer()
	defer shutdown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	opts := NewOptions("private-key")
	opts.BaseURL = server.URL
	client := NewClientWithOptions(opts)

	ok, err := client.Confirm("126.0.0.1", "any-recaptcha-response")
	expectedError := "Unexpected response. Expected 200 but found 400"
	if err.Error() != expectedError {
		t.Errorf("Unexpected error: %s", err)
	}

	if ok {
		t.Error("Expected ok to be false")
	}
}

func TestInvalidJSONUsingMockServer(t *testing.T) {
	mux, server, shutdown := setupMockServer()
	defer shutdown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "invalid JSON")
	})

	opts := NewOptions("private-key")
	opts.BaseURL = server.URL
	client := NewClientWithOptions(opts)

	ok, err := client.Confirm("126.0.0.1", "any-recaptcha-response")
	expectedError := "invalid character 'i' looking for beginning of value"
	if err.Error() != expectedError {
		t.Errorf("Unexpected error: %s", err)
	}

	if ok {
		t.Error("Expected ok to be false")
	}
}

func TestConfirmUsingMockServer(t *testing.T) {
	mux, server, shutdown := setupMockServer()
	defer shutdown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"success": true,
				"challenge_ts": "2018-05-28T15:30:00Z",
				"hostname": "localhost",
				"error_codes": []
			}`)
	})

	opts := NewOptions("private-key")
	opts.BaseURL = server.URL
	client := NewClientWithOptions(opts)

	ok, err := client.Confirm("126.0.0.1", "any-recaptcha-response")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !ok {
		t.Error("Expected ok to be true")
	}
}
