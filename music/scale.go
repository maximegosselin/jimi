package music

import (
	"strconv"
	"strings"
)

var patterns = map[string][]int{
	"aeolian":          {2, 1, 2, 2, 1, 2, 2},
	"chromatic":        {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	"dorian":           {2, 1, 2, 2, 2, 1, 2},
	"ionian":           {2, 2, 1, 2, 2, 2, 1},
	"locrian":          {1, 2, 2, 1, 2, 2, 2},
	"lydian":           {2, 2, 2, 1, 2, 2, 1},
	"major":            {2, 2, 1, 2, 2, 2, 1},
	"major-blues":      {2, 1, 1, 3, 2, 3},
	"major-pentatonic": {2, 2, 3, 2, 3},
	"minor":            {2, 1, 2, 2, 1, 2, 2},
	"minor-blues":      {3, 2, 1, 1, 3, 2},
	"minor-pentatonic": {3, 2, 2, 3, 2},
	"myxolydian":       {2, 2, 1, 2, 2, 1, 2},
	"phrygian":         {1, 2, 2, 2, 1, 2, 2},

	"major-triad":  {4, 3},
	"minor-triad":  {3, 4},
	"sus2":         {2, 5},
	"sus4":         {5, 2},
	"aug":          {4, 4},
	"dim":          {3, 3},
	"5th":          {7},
	"major-6th":    {4, 3, 2},
	"minor-6th":    {3, 4, 2},
	"dominant-7th": {4, 3, 3},
	"major-7th":    {4, 3, 4},
	"minor-7th":    {3, 4, 3},
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

func NewScale(pattern string, key *Note) *Scale {
	pattern = strings.ToLower(pattern)
	scale := Scale{pattern, key, make(map[string]*Note)}
	if intervals, ok := patterns[pattern]; ok {
		scale.resolveNotes(key, intervals)
		return &scale
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
	return &scale
}
