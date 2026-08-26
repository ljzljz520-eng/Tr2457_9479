package archive

import (
	"frontend_go/model"
	"sort"
)

func Archived(rs []model.Record) []model.Record {
	out := make([]model.Record, 0)
	for _, r := range rs {
		if r.IsArchived() {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.Before(out[j].UpdatedAt) })
	return out
}
func CountByPublisher(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		if r.IsArchived() {
			m[r.Publisher]++
		}
	}
	return m
}
