package store

import (
	"fmt"
	"frontend_go/model"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var bucketRecords = []byte("records")
var bucketProfiles = []byte("profiles")
var bucketEvents = []byte("events")
var bucketAudits = []byte("audits")

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range [][]byte{bucketRecords, bucketProfiles, bucketEvents, bucketAudits} {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func (s *Store) Path() string                   { return s.path }
func (s *Store) PutRecord(r model.Record) error { return s.put(bucketRecords, r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.get(bucketRecords, id, &r)
	if e != nil {
		return r, e
	}
	return r, nil
}
func (s *Store) DeleteRecord(id string) error { return s.del(bucketRecords, id) }
func (s *Store) ListRecords() ([]model.Record, error) {
	var out []model.Record
	e := s.list(bucketRecords, &out)
	return out, e
}
func (s *Store) PutProfile(p model.Profile) error { return s.put(bucketProfiles, p.ID, p) }
func (s *Store) GetProfile(id string) (model.Profile, error) {
	var p model.Profile
	e := s.get(bucketProfiles, id, &p)
	return p, e
}
func (s *Store) PutEvent(v model.Event) error { return s.put(bucketEvents, v.ID, v) }
func (s *Store) GetEvent(id string) (model.Event, error) {
	var v model.Event
	e := s.get(bucketEvents, id, &v)
	return v, e
}
func (s *Store) PutAudit(a model.Audit) error { return s.put(bucketAudits, a.ID, a) }
func (s *Store) ListAudits() ([]model.Audit, error) {
	var a []model.Audit
	e := s.list(bucketAudits, &a)
	return a, e
}
func (s *Store) put(bucket []byte, id string, v any) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}
	data, e := model.Encode(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(id), data) })
}
func (s *Store) get(bucket []byte, id string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		d := tx.Bucket(bucket).Get([]byte(id))
		if d == nil {
			return model.ErrNotFound
		}
		return model.Decode(d, v)
	})
}
func (s *Store) del(bucket []byte, id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Delete([]byte(id)) })
}
func (s *Store) list(bucket []byte, out any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		var raw [][]byte
		e := tx.Bucket(bucket).ForEach(func(_, v []byte) error { raw = append(raw, append([]byte(nil), v...)); return nil })
		if e != nil {
			return e
		}
		switch x := out.(type) {
		case *[]model.Record:
			for _, d := range raw {
				var r model.Record
				if e := model.Decode(d, &r); e != nil {
					return e
				}
				*x = append(*x, r)
			}
		case *[]model.Audit:
			for _, d := range raw {
				var a model.Audit
				if e := model.Decode(d, &a); e != nil {
					return e
				}
				*x = append(*x, a)
			}
		}
		return nil
	})
}
