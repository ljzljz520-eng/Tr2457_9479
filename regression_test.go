package frontend_go

import (
	"context"
	"frontend_go/events"
	"frontend_go/model"
	"frontend_go/store"
	"os"
	"path/filepath"
	"testing"
)

func TestBusinessChain31(t *testing.T) {
	if os.Getenv("RUN_BUG_REGRESSION") == "" {
		t.Skip("known injected regression")
	}
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	d := events.NewDispatcher(s, nil)
	if err := d.Publish(context.Background(), model.Event{ID: "bug31", Kind: "published"}); err != nil {
		t.Fatalf("optional auditor should not block publish: %v", err)
	}
}
