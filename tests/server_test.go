package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/sabariramc/WordCURD/pkg/app"
	"github.com/sabariramc/WordCURD/pkg/utils"
	"gotest.tools/assert"
)

func TestServerRequest(t *testing.T) {
	srv, err := app.NewApp()
	assert.NilError(t, err)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	statusCode := w.Result().StatusCode
	assert.Equal(t, statusCode, http.StatusOK)
	res := string(w.Body.Bytes())
	assert.Equal(t, res, "Hello, World!")
}

func TestWordAddition(t *testing.T) {
	srv, err := app.NewApp()
	assert.NilError(t, err)
	newWord := fmt.Sprintf("test-%v", uuid.New().String())
	body := map[string]string{
		"Word": newWord,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(body)
	assert.NilError(t, err)
	req := httptest.NewRequest("POST", "/word", &buf)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	statusCode := w.Result().StatusCode
	assert.Equal(t, statusCode, http.StatusOK)
	srv, err = app.NewApp()
	assert.NilError(t, err)
	req = httptest.NewRequest("GET", "/word", &buf)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	statusCode = w.Result().StatusCode
	assert.Equal(t, statusCode, http.StatusOK)
	wordList := make([]string, 0)
	err = json.NewDecoder(w.Body).Decode(&wordList)
	assert.NilError(t, err, "Get response read error")
	found := false
	for _, w := range wordList {
		if w == newWord {
			found = true
		}
	}
	assert.Equal(t, found, true, "Word not found")
}
