package model

import "time"

type Record struct {
	ID, Title, Publisher, Status string
	Rating                       float64
	Tags                         []string
	CreatedAt, UpdatedAt         time.Time
}
type Profile struct {
	ID, Name, Email string
	FavoriteTags    []string
	CreatedAt       time.Time
}
type Event struct {
	ID, RecordID, Kind, Payload string
	CreatedAt                   time.Time
	Delivered                   bool
}
type Audit struct {
	ID, EventID, Action, Detail string
	CreatedAt                   time.Time
}

func NewRecord(id, title, publisher string) Record {
	now := time.Now().UTC()
	return Record{ID: id, Title: title, Publisher: publisher, Status: "new", CreatedAt: now, UpdatedAt: now}
}
func (r Record) Valid() bool { return r.ID != "" && r.Title != "" && r.Rating >= 0 && r.Rating <= 10 }
func (r *Record) Touch()     { r.UpdatedAt = time.Now().UTC() }
func (r *Record) SetRating(v float64) error {
	if v < 0 || v > 10 {
		return ErrRating
	}
	r.Rating = v
	r.Touch()
	return nil
}
func (r *Record) AddTag(tag string) {
	if tag != "" {
		for _, t := range r.Tags {
			if t == tag {
				return
			}
		}
		r.Tags = append(r.Tags, tag)
		r.Touch()
	}
}
func (r *Record) Archive()        { r.Status = "archived"; r.Touch() }
func (r Record) IsArchived() bool { return r.Status == "archived" }

type sentinel string

func (e sentinel) Error() string { return string(e) }

var ErrRating sentinel = "rating must be between 0 and 10"
var ErrNotFound sentinel = "record not found"
