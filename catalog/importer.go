package catalog

import (
	"bufio"
	"frontend_go/model"
	"io"
	"strings"
)

func ImportLines(s *Service, r io.Reader) (int, error) {
	scan := bufio.NewScanner(r)
	count := 0
	for scan.Scan() {
		parts := strings.Split(scan.Text(), "|")
		if len(parts) < 3 {
			continue
		}
		rec := model.NewRecord(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2]))
		if len(parts) > 3 {
			rec.Tags = model.NormalizeTags(strings.Split(parts[3], ","))
		}
		if err := s.Register(rec); err != nil {
			return count, err
		}
		count++
	}
	return count, scan.Err()
}
func ExportLines(rs []model.Record) string {
	var b strings.Builder
	for _, r := range rs {
		b.WriteString(r.ID)
		b.WriteByte('|')
		b.WriteString(r.Title)
		b.WriteByte('|')
		b.WriteString(r.Publisher)
		b.WriteByte('|')
		b.WriteString(strings.Join(r.Tags, ","))
		b.WriteByte('\n')
	}
	return b.String()
}
