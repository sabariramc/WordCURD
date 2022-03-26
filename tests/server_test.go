package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sabariramc/WordCURD/pkg/app"
	"github.com/sabariramc/WordCURD/pkg/utils"
)

func TestServerRequest(t *testing.T) {
	srv, err := app.NewApp()
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	statusCode := w.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Fatalf("Status Code %v", statusCode)
	}
}

func TestWordAddition(t *testing.T) {
	srv, err := app.NewApp()
	if err != nil {
		t.Fatal(err)
	}
	body := map[string]string{
		"Word": "test",
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("GET", "/word", &buf)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	statusCode := w.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Fatalf("Status Code %v", statusCode)
	}
}
