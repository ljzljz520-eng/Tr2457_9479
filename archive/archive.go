package archive

import (
	"fmt"
	"frontend_go/catalog"
	"frontend_go/model"
	"frontend_go/store"
	"time"
)

type Manager struct {
	catalog *catalog.Service
	db      *store.Store
}

func New(c *catalog.Service, d *store.Store) *Manager { return &Manager{catalog: c, db: d} }
func (m *Manager) Archive(id string) error            { return m.catalog.Archive(id) }
func (m *Manager) Restore(id string) error {
	r, e := m.db.GetRecord(id)
	if e != nil {
		return e
	}
	if !r.IsArchived() {
		return fmt.Errorf("record is active")
	}
	r.Status = "restored"
	r.Touch()
	return m.db.PutRecord(r)
}
func (m *Manager) Snapshot(id string) (model.Record, error) {
	r, e := m.db.GetRecord(id)
	if e != nil {
		return r, e
	}
	r.UpdatedAt = time.Now().UTC()
	return model.CloneRecord(r), nil
}
