package events

import (
	"frontend_go/model"
	"sync"
)

type Queue struct {
	mu    sync.Mutex
	items []model.Event
}

func (q *Queue) Push(e model.Event) { q.mu.Lock(); defer q.mu.Unlock(); q.items = append(q.items, e) }
func (q *Queue) Pop() (model.Event, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return model.Event{}, false
	}
	e := q.items[0]
	q.items = q.items[1:]
	return e, true
}
func (q *Queue) Len() int { q.mu.Lock(); defer q.mu.Unlock(); return len(q.items) }
func (q *Queue) Drain() []model.Event {
	q.mu.Lock()
	defer q.mu.Unlock()
	out := append([]model.Event(nil), q.items...)
	q.items = nil
	return out
}
