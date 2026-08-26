package scoring

import (
	"frontend_go/model"
	"sort"
)

type Forecast struct {
	RecordID   string
	Value      float64
	Confidence float64
}

func (e *Engine) Forecast(r model.Record, history []float64) Forecast {
	if len(history) == 0 {
		return Forecast{RecordID: r.ID, Value: r.Rating, Confidence: .2}
	}
	avg := WeightedAverage(history, nil)
	confidence := float64(len(history)) / 10
	if confidence > 1 {
		confidence = 1
	}
	return Forecast{RecordID: r.ID, Value: (avg + r.Rating) / 2, Confidence: confidence}
}
func Normalize(rs []model.Record) []model.Record {
	out := append([]model.Record(nil), rs...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Rating > out[j].Rating })
	return out
}
