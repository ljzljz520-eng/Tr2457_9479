package catalog

import (
	"frontend_go/model"
	"strings"
)

func Match(r model.Record, q string) bool {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return true
	}
	for _, t := range r.Tags {
		if strings.Contains(strings.ToLower(t), q) {
			return true
		}
	}
	return strings.Contains(strings.ToLower(r.Title), q) || strings.Contains(strings.ToLower(r.Publisher), q)
}
func GroupByStatus(rs []model.Record) map[string][]model.Record {
	m := map[string][]model.Record{}
	for _, r := range rs {
		m[r.Status] = append(m[r.Status], r)
	}
	return m
}
