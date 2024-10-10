package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/* Integration Testing */
func TestServerAvailability(t *testing.T) {
	check := t.Run("Check if port is open", func(t *testing.T) {
		timeout := time.Second
		port := API_PORT
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
		requestUrl := fmt.Sprintf("http://localhost:%d/%s", API_PORT, API_PATH)
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

/* Unit Testing */
func TestCheckHttpHeaders(t *testing.T) {

	assertHttpStatus := func(got, expected int) {
		if got != expected {
			t.Errorf("Got %d , Expected %d", got, expected)
		}
	}

	createRequestResponse :=
		func() (*http.Request, *httptest.ResponseRecorder) {
			request, _ := http.NewRequest(http.MethodGet, "/auth", nil)
			response := httptest.NewRecorder()

			return request, response
		}

	t.Run("Expect BadRequest status , when ip and username is not set", func(t *testing.T) {

		request, response := createRequestResponse()
		request.Header.Set("X-Original-IP", "")
		request.Header.Set("X-Username", "")

		AuthHandler(response, request)
		got := response.Result().StatusCode
		expected := http.StatusBadRequest

		assertHttpStatus(got, expected)
	})

	t.Run("Expect OK status , when both ip & username is set", func(t *testing.T) {
		request, response := createRequestResponse()
		request.Header.Set("X-Original-IP", "127.0.0.1")
		request.Header.Set("X-Username", "123456")

		AuthHandler(response, request)
		got := response.Result().StatusCode
		expected := http.StatusOK

		assertHttpStatus(got, expected)
	})
}
