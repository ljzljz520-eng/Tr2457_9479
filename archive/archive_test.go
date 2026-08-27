package archive

import (
	"frontend_go/catalog"
	"frontend_go/model"
	"frontend_go/store"
	"path/filepath"
	"testing"
)

func TestArchiveRestore(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := catalog.New(s)
	_ = c.Register(model.NewRecord("1", "Game", "Pub"))
	m := New(c, s)
	if e := m.Archive("1"); e != nil {
		t.Fatal(e)
	}
	if e := m.Restore("1"); e != nil {
		t.Fatal(e)
	}
}
