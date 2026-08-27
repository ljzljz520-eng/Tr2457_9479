package model

import "strings"

func (p Profile) Valid() bool { return p.ID != "" && strings.TrimSpace(p.Name) != "" }
func (p *Profile) AddFavorite(tag string) {
	tag = strings.TrimSpace(strings.ToLower(tag))
	if tag == "" {
		return
	}
	for _, x := range p.FavoriteTags {
		if x == tag {
			return
		}
	}
	p.FavoriteTags = append(p.FavoriteTags, tag)
}
func (p Profile) Likes(r Record) bool {
	for _, want := range p.FavoriteTags {
		for _, tag := range r.Tags {
			if want == tag {
				return true
			}
		}
	}
	return false
}
