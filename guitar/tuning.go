package guitar

import (
	"errors"
	"github.com/maximegosselin/jimi/music"
	"regexp"
)

type Tuning [6]*music.Note

func NewTuning(t string) (Tuning, error) {
	re := regexp.MustCompile("([ABCDEFG][#b]?)")
	matches := re.FindAllString(t, -1)
	if len(matches) != 6 {
		return Tuning{}, errors.New("invalid tuning")
	}

	var tuning Tuning
	for i, match := range matches {
		tuning[i], _ = music.NewNote(match)
	}
	return tuning, nil
}
