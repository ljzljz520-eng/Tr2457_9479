package scoring

import "frontend_go/model"

func Clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 10 {
		return 10
	}
	return v
}
func WeightedAverage(values []float64, weights []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var s, w float64
	for i, v := range values {
		weight := 1.0
		if i < len(weights) {
			weight = weights[i]
		}
		s += Clamp(v) * weight
		w += weight
	}
	if w == 0 {
		return 0
	}
	return s / w
}
func Eligible(r model.Record) bool { return !r.IsArchived() && r.Rating > 0 }
