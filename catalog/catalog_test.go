package catalog

import (
	"frontend_go/model"
	"frontend_go/store"
	"path/filepath"
	"testing"
)

func TestCatalogSearch(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(s)
	_ = c.Register(model.NewRecord("1", "Azul", "Plan B"))
	r, e := c.Search("azul")
	if e != nil || len(r) != 1 {
		t.Fatal(e)
	}
}
