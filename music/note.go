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
	n = normalizeNote(n)
	re := regexp.MustCompile("([ABCDEFG][#]?)")
	match := re.FindString(n)
	if match != n {
		return nil, errors.New("invalid note")
	}
	return chromatic[n], nil
}

func normalizeNote(n string) string {
	switch n {
	case "Cb":
		return "B"
	case "Db":
		return "C#"
	case "Eb":
		return "D#"
	case "E#":
		return "F"
	case "Fb":
		return "E"
	case "Gb":
		return "F#"
	case "Ab":
		return "G#"
	case "Bb":
		return "A#"
	case "B#":
		return "C"
	default:
		return n
	}
}
