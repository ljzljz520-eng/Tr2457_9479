package events

import (
	"context"
	"frontend_go/model"
	"frontend_go/store"
	"path/filepath"
	"testing"
)

type testAuditor struct{ called bool }

func (a *testAuditor) Audit(context.Context, model.Event) error { a.called = true; return nil }
func TestDispatcherWithAuditor(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	a := &testAuditor{}
	d := NewDispatcher(s, a)
	if e := d.Publish(context.Background(), model.Event{ID: "e1", Kind: "added"}); e != nil || !a.called {
		t.Fatal(e)
	}
}
