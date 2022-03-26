package tests

import (
	"bytes"
	"encoding/json"
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
	body := map[string]string{
		"Word": "test",
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req := httptest.NewRequest("POST", "/tenant", &buf)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
}
