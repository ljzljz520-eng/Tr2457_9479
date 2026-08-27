package model

import "testing"

func TestRecordValidation(t *testing.T) {
	r := NewRecord("1", "Catan", "Kosmos")
	if !r.Valid() {
		t.Fatal("valid record rejected")
	}
	if e := r.SetRating(11); e == nil {
		t.Fatal("invalid rating accepted")
	}
}
