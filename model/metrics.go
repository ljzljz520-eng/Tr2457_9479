package model

import "time"

type CollectionMetrics struct {
	Total, Archived, Rated int
	Average                float64
	Latest                 time.Time
}

func BuildMetrics(rs []Record) CollectionMetrics {
	m := CollectionMetrics{}
	for _, r := range rs {
		m.Total++
		if r.IsArchived() {
			m.Archived++
		}
		if r.Rating > 0 {
			m.Rated++
			m.Average += r.Rating
		}
		if r.UpdatedAt.After(m.Latest) {
			m.Latest = r.UpdatedAt
		}
	}
	if m.Rated > 0 {
		m.Average /= float64(m.Rated)
	}
	return m
}
func (m CollectionMetrics) Completion() float64 {
	if m.Total == 0 {
		return 0
	}
	return float64(m.Rated) / float64(m.Total)
}
func (m CollectionMetrics) Healthy() bool { return m.Total > 0 && m.Average >= 5 }
