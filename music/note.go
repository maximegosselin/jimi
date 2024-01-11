package music

import (
	"errors"
	"regexp"
)

type Note struct {
	name  string
	color string
	next  *Note
}

func (n Note) Name() string {
	return n.name
}

func (n Note) Color() string {
	return n.color
}

func (n Note) Next() *Note {
	return n.next
}

func (n Note) Equals(other *Note) bool {
	return n.name == other.name
}

func (n Note) String() string {
	return n.name
}

func NewNote(n string) (*Note, error) {
	n = normalizeName(n)
	re := regexp.MustCompile("([ABCDEFG][#]?)")
	match := re.FindString(n)
	if match != n {
		return nil, errors.New("invalid note")
	}
	return chromatic[n], nil
}
