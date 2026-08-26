package scoring

import (
	"frontend_go/model"
	"testing"
)

func TestScoring(t *testing.T) {
	e := New()
	r := model.NewRecord("1", "A", "B")
	r.Rating = 8
	if e.Explain(r) != "highly recommended" {
		t.Fatal("explain")
	}
	if WeightedAverage([]float64{8, 6}, nil) != 7 {
		t.Fatal("average")
	}
}
