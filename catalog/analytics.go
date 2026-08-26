package catalog

import (
	"frontend_go/model"
	"sort"
	"strings"
	"time"
)

type TrendPoint struct {
	Day     time.Time
	Count   int
	Average float64
}

func Titles(rs []model.Record) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.Title)
	}
	sort.Strings(out)
	return out
}
func Publishers(rs []model.Record) []string {
	m := map[string]bool{}
	for _, r := range rs {
		m[r.Publisher] = true
	}
	out := make([]string, 0, len(m))
	for p := range m {
		out = append(out, p)
	}
	sort.Strings(out)
	return out
}
func Tags(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		for _, t := range r.Tags {
			m[strings.ToLower(t)]++
		}
	}
	return m
}
func TopRated(rs []model.Record, n int) []model.Record {
	if n < 0 {
		n = 0
	}
	out := append([]model.Record(nil), rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].Rating > out[j].Rating })
	if n < len(out) {
		out = out[:n]
	}
	return out
}
func RecentlyUpdated(rs []model.Record, since time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.UpdatedAt.After(since) {
			out = append(out, r)
		}
	}
	return out
}
func RatingBuckets(rs []model.Record) map[int]int {
	m := map[int]int{}
	for _, r := range rs {
		bucket := int(r.Rating)
		if bucket < 0 {
			bucket = 0
		}
		if bucket > 10 {
			bucket = 10
		}
		m[bucket]++
	}
	return m
}
func StatusCounts(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func HasPublisher(rs []model.Record, p string) bool {
	for _, r := range rs {
		if strings.EqualFold(r.Publisher, p) {
			return true
		}
	}
	return false
}
func FindByID(rs []model.Record, id string) (model.Record, bool) {
	for _, r := range rs {
		if r.ID == id {
			return r, true
		}
	}
	return model.Record{}, false
}
func MergeTags(a, b []string) []string {
	return model.NormalizeTags(append(append([]string{}, a...), b...))
}
func FilterRating(rs []model.Record, min, max float64) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.Rating >= min && r.Rating <= max {
			out = append(out, r)
		}
	}
	return out
}
func FilterTag(rs []model.Record, tag string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		for _, t := range r.Tags {
			if strings.EqualFold(t, tag) {
				out = append(out, r)
				break
			}
		}
	}
	return out
}
func StableSortTitle(rs []model.Record) []model.Record {
	out := append([]model.Record(nil), rs...)
	sort.SliceStable(out, func(i, j int) bool { return strings.ToLower(out[i].Title) < strings.ToLower(out[j].Title) })
	return out
}
func DistinctTitles(rs []model.Record) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, r := range rs {
		key := strings.ToLower(r.Title)
		if !seen[key] {
			seen[key] = true
			out = append(out, r.Title)
		}
	}
	return out
}
func AverageForPublisher(rs []model.Record) map[string]float64 {
	sum := map[string]float64{}
	count := map[string]int{}
	for _, r := range rs {
		sum[r.Publisher] += r.Rating
		count[r.Publisher]++
	}
	out := map[string]float64{}
	for p, v := range sum {
		out[p] = v / float64(count[p])
	}
	return out
}
func Timeline(rs []model.Record) []TrendPoint {
	days := map[string]*TrendPoint{}
	for _, r := range rs {
		k := r.CreatedAt.Format("2006-01-02")
		p := days[k]
		if p == nil {
			d, _ := time.Parse("2006-01-02", k)
			p = &TrendPoint{Day: d}
			days[k] = p
		}
		p.Count++
		p.Average += r.Rating
	}
	out := make([]TrendPoint, 0, len(days))
	for _, p := range days {
		p.Average /= float64(p.Count)
		out = append(out, *p)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Day.Before(out[j].Day) })
	return out
}
func ValidateBatch(rs []model.Record) []string {
	errs := []string{}
	for _, r := range rs {
		if !r.Valid() {
			errs = append(errs, r.ID)
		}
	}
	return errs
}
func ApplyTag(rs []model.Record, tag string) []model.Record {
	out := make([]model.Record, len(rs))
	for i, r := range rs {
		r.AddTag(tag)
		out[i] = r
	}
	return out
}
func CloneAll(rs []model.Record) []model.Record {
	out := make([]model.Record, len(rs))
	for i, r := range rs {
		out[i] = model.CloneRecord(r)
	}
	return out
}
func Empty(rs []model.Record) bool { return len(rs) == 0 }
