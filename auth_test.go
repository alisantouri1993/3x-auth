package main

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

/* Integration Testing */
func TestServerAvailability(t *testing.T) {
	check := t.Run("Check if port is open", func(t *testing.T) {
		timeout := time.Second
		port := 20000
		address := fmt.Sprintf("localhost:%d", port)
		_, err := net.DialTimeout("tcp", address, timeout)

		if err != nil {
			t.Fatalf("Error Opening Port %d , ERROR= %q", port, err)
		}
	})

	if !check {
		t.Fatal("Check Error!")
	}

	t.Run("Test Auth API response is OK", func(t *testing.T) {
		requestUrl := fmt.Sprintf("http://localhost:%d/%s", 20000, "/auth")
		req, err := http.NewRequest(http.MethodGet, requestUrl, nil)

		if err != nil {
			t.Errorf("Error Creating Request , ERROR= %q", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("Error Sending request , ERROR= %q", err)
		}

		if resp != nil && resp.StatusCode != http.StatusOK {
			t.Errorf("Request not Accepted , response code = %d", resp.StatusCode)
		}
	})
}
