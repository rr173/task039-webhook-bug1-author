package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"task039-webhook/internal/webhook"
)

func TestProbeEmptyEventListIsJSONArray(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/events", nil)
	rec := httptest.NewRecorder()
	New(webhook.New("probe-secret")).Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d want 200", rec.Code)
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte(`"events":[]`)) {
		t.Fatalf("empty event list must be JSON array, body=%s", rec.Body.Bytes())
	}
}
