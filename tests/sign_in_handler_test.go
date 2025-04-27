package tests

import (
	"bytes"
	"c2nofficialsitebackend/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignInHandler(t *testing.T) {

	dummyData := map[string]string{
		"email":     "harshsason2000@gmail.com",
		"password":  "HarhsMohan55#",
		"auth_type": "email",
	}
	body, _ := json.Marshal(dummyData)

	req := httptest.NewRequest("POST", "/signin", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Control-Allow-Credentials", "true")

	rr := httptest.NewRecorder()

	handlers.ReceiveSignInFormUserInfo(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", rr.Code)
	}
}
