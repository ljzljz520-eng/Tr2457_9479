package model

import (
	"encoding/json"
	"strings"
)

func Encode(v any) ([]byte, error)    { return json.Marshal(v) }
func Decode(data []byte, v any) error { return json.Unmarshal(data, v) }
func NormalizeTags(tags []string) []string {
	out := make([]string, 0, len(tags))
	seen := map[string]bool{}
	for _, t := range tags {
		t = strings.TrimSpace(strings.ToLower(t))
		if t != "" && !seen[t] {
			seen[t] = true
			out = append(out, t)
		}
	}
	return out
}
func CloneRecord(r Record) Record { r.Tags = append([]string(nil), r.Tags...); return r }
