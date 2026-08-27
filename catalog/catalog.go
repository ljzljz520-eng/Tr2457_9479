package catalog

import (
	"fmt"
	"frontend_go/model"
	"frontend_go/store"
	"sort"
	"strings"
)

type Service struct{ db *store.Store }

func New(db *store.Store) *Service { return &Service{db: db} }
func (s *Service) Register(r model.Record) error {
	if !r.Valid() {
		return fmt.Errorf("invalid record")
	}
	r.Tags = model.NormalizeTags(r.Tags)
	return s.db.PutRecord(r)
}
func (s *Service) UpdateRating(id string, v float64) error {
	r, e := s.db.GetRecord(id)
	if e != nil {
		return e
	}
	if e = r.SetRating(v); e != nil {
		return e
	}
	return s.db.PutRecord(r)
}
func (s *Service) AddTag(id, tag string) error {
	r, e := s.db.GetRecord(id)
	if e != nil {
		return e
	}
	r.AddTag(tag)
	return s.db.PutRecord(r)
}
func (s *Service) Search(query string) ([]model.Record, error) {
	all, e := s.db.ListRecords()
	if e != nil {
		return nil, e
	}
	q := strings.ToLower(strings.TrimSpace(query))
	out := all
	if q != "" {
		out = nil
		for _, r := range all {
			if strings.Contains(strings.ToLower(r.Title), q) || strings.Contains(strings.ToLower(r.Publisher), q) {
				out = append(out, r)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	return out, nil
}
func (s *Service) Archive(id string) error {
	r, e := s.db.GetRecord(id)
	if e != nil {
		return e
	}
	r.Archive()
	return s.db.PutRecord(r)
}
func (s *Service) Summary() (int, float64, error) {
	rs, e := s.db.ListRecords()
	if e != nil {
		return 0, 0, e
	}
	var total float64
	for _, r := range rs {
		total += r.Rating
	}
	if len(rs) == 0 {
		return 0, 0, nil
	}
	return len(rs), total / float64(len(rs)), nil
}
