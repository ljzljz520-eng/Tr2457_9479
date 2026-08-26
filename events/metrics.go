package events

import (
	"frontend_go/model"
	"sync/atomic"
)

type Metrics struct {
	published atomic.Int64
	failed    atomic.Int64
}

func (m *Metrics) Published()               { m.published.Add(1) }
func (m *Metrics) Failed()                  { m.failed.Add(1) }
func (m *Metrics) Snapshot() (int64, int64) { return m.published.Load(), m.failed.Load() }
func EventSummary(es []model.Event) (int, int) {
	delivered := 0
	for _, e := range es {
		if e.Delivered {
			delivered++
		}
	}
	return len(es), delivered
}
