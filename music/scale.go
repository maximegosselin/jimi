package music

import (
	"strings"
)

type Scale struct {
	name  string
	root  *Note
	notes map[string]*Note
}

func (s Scale) Name() string {
	return s.name
}

func (s Scale) Root() *Note {
	return s.root
}

func (s Scale) ContainsNote(n *Note) bool {
	_, found := s.notes[n.Name()]
	return found
}

func (s Scale) RootIs(n *Note) bool {
	return s.root.Equals(n)
}

func (s Scale) resolveNotes(root *Note, intervals []string) {
	for _, i := range intervals {
		next := root
		interval, _ := NewInterval(i)
		for semitones := 0; semitones < interval.Semitones(); semitones++ {
			next = next.Next()
		}
		s.notes[next.Name()] = next
	}
}

func NewScale(pattern string, root *Note) *Scale {
	scale := Scale{pattern, root, make(map[string]*Note)}
	if intervals, ok := patterns[pattern]; ok {
		scale.resolveNotes(root, intervals)
		return &scale
	}

	var intervals []string
	parts := strings.Split(pattern, "-")
	for _, p := range parts {
		if validInterval(p) {
			intervals = append(intervals, p)
		}
	}

	scale.resolveNotes(root, intervals)
	return &scale
}
