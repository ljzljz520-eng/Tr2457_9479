package events

import (
	"context"
	"fmt"
	"frontend_go/model"
	"frontend_go/store"
	"time"
)

type Auditor interface {
	Audit(context.Context, model.Event) error
}
type AuditSink struct{ db *store.Store }

func NewAuditSink(db *store.Store) *AuditSink { return &AuditSink{db: db} }
func (a *AuditSink) Audit(ctx context.Context, e model.Event) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return a.db.PutAudit(model.Audit{ID: e.ID + "-audit", EventID: e.ID, Action: e.Kind, Detail: e.Payload, CreatedAt: time.Now().UTC()})
}

type Dispatcher struct {
	db      *store.Store
	auditor Auditor
}

func NewDispatcher(db *store.Store, a Auditor) *Dispatcher { return &Dispatcher{db: db, auditor: a} }
func (d *Dispatcher) Publish(ctx context.Context, e model.Event) error {
	if e.ID == "" {
		return fmt.Errorf("event id required")
	}
	if err := d.db.PutEvent(e); err != nil {
		return err
	}
	if err := d.auditor.Audit(ctx, e); err != nil {
		return err
	}
	e.Delivered = true
	return d.db.PutEvent(e)
}
func (d *Dispatcher) Replay(ctx context.Context, id string) error {
	e, err := d.db.GetEvent(id)
	if err != nil {
		return err
	}
	if e.Delivered {
		return nil
	}
	return d.Publish(ctx, e)
}
func (d *Dispatcher) Pending() ([]model.Event, error) { return nil, nil }
