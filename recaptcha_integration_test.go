package recaptcha

import "testing"

func TestConfirm(t *testing.T) {
	// If run in short mode, we want to skip integration testing, which requires
	// network requests.
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	// Create a new client using the private key that Google has defined for
	// testing.
	client := NewClient("6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe")

	// The test key will always allow verification requests to pass, so we can use
	// any recaptcha response.
	ok, err := client.Confirm("126.0.0.1", "any-recaptcha-response")
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Error("Response from Recaptcha service didn't verify response for client")
	}
}
