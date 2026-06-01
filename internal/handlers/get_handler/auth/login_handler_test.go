package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	handler := Auth{}

	handler.LoginHandler(recorder, req)

	if recorder.Code == http.StatusOK {
		body := recorder.Body
		t.Log(body.String())
	} else {
		body := recorder.Body
		t.Fatal("Unexpected response: ", body.String())
	}
}
