package store

import (
	"frontend_go/model"
)

func AverageRating(rs []model.Record) float64 {
	if len(rs) == 0 {
		return 0
	}
	var s float64
	for _, r := range rs {
		s += r.Rating
	}
	return s / float64(len(rs))
}
func FilterArchived(rs []model.Record, want bool) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.IsArchived() == want {
			out = append(out, r)
		}
	}
	return out
}
