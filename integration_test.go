package frontend_go

import (
	"context"
	"frontend_go/catalog"
	"frontend_go/events"
	"frontend_go/model"
	"frontend_go/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := catalog.New(s)
	r := model.NewRecord("w1", "Ticket to Ride", "Days")
	if e := c.Register(r); e != nil {
		t.Fatal(e)
	}
	found, e := c.Search("ticket")
	if e != nil || len(found) != 1 {
		t.Fatal("not found")
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := catalog.New(s)
	r := model.NewRecord("w2", "Root", "Leder")
	_ = c.Register(r)
	_ = c.UpdateRating("w2", 8)
	_ = c.Archive("w2")
	got, _ := s.GetRecord("w2")
	if !got.IsArchived() {
		t.Fatal("not archived")
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	d := events.NewDispatcher(s, &noopAuditor{})
	e := model.Event{ID: "w3", Kind: "published"}
	if err := d.Publish(context.Background(), e); err != nil {
		t.Fatal(err)
	}
}

type noopAuditor struct{}

func (*noopAuditor) Audit(context.Context, model.Event) error { return nil }
