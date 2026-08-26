package server

import (
	"frontend_go/catalog"
	"frontend_go/store"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHTTPRecords(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	h := New(catalog.New(s), nil).Handler()
	r := httptest.NewRequest("GET", "/records", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatalf("status %d", w.Code)
	}
}
