package scoring

import (
	"frontend_go/model"
	"math"
)

type Engine struct{ weights map[string]float64 }

func New() *Engine {
	return &Engine{weights: map[string]float64{"complexity": .35, "theme": .35, "replay": .3}}
}
func (e *Engine) Score(r model.Record, factors map[string]float64) float64 {
	v := r.Rating
	for k, w := range e.weights {
		if x, ok := factors[k]; ok {
			v = v*(1-w) + x*w
		}
	}
	return math.Round(v*100) / 100
}
func (e *Engine) Rank(rs []model.Record) []model.Record {
	out := append([]model.Record(nil), rs...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Rating > out[i].Rating {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
func (e *Engine) Explain(r model.Record) string {
	if r.Rating >= 8 {
		return "highly recommended"
	}
	if r.Rating >= 5 {
		return "worth exploring"
	}
	return "niche appeal"
}
