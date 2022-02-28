package music

import (
	"strconv"
	"strings"
)

var patterns = map[string][]int{
	"chromatic":        {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	"major":            {2, 2, 1, 2, 2, 2, 1},
	"pentatonic":       {2, 2, 3, 2, 3},
	"minor":            {2, 1, 2, 2, 1, 2, 2},
	"minor-pentatonic": {3, 2, 2, 3, 2},

	"triad":       {4, 3},
	"minor-triad": {3, 4},
	"sus2":        {2, 5},
	"sus4":        {5, 2},
	"dim":         {3, 3},
	"fifth":       {7},
	"major-7th":   {4, 3, 4},
	"minor-7th":   {3, 4, 3},
}

type Scale struct {
	name  string
	key   *Note
	notes map[string]*Note
}

func (s Scale) Name() string {
	return s.name
}

func (s Scale) Key() *Note {
	return s.key
}

func (s Scale) ContainsNote(n *Note) bool {
	_, found := s.notes[n.Name()]
	return found
}

func (s Scale) KeyIs(n *Note) bool {
	return s.key.Equals(n)
}

func (s Scale) resolveNotes(key *Note, intervals []int) {
	s.notes[key.Name()] = key
	next := key
	for _, step := range intervals {
		for i := 0; i < step; i++ {
			next = next.Next()
		}
		s.notes[next.Name()] = next
	}
}

func NewScale(pattern string, key *Note) (*Scale, error) {
	pattern = strings.ToLower(pattern)
	scale := Scale{pattern, key, make(map[string]*Note)}
	if intervals, ok := patterns[pattern]; ok {
		scale.resolveNotes(key, intervals)
		return &scale, nil
	}

	var intervals []int
	parts := strings.Split(pattern, "-")
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err == nil && i > 0 && i < 12 {
			intervals = append(intervals, i)
		}
	}

	scale.resolveNotes(key, intervals)
	return &scale, nil
}
