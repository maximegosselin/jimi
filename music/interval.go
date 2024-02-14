package music

import (
	"errors"
)

var semitones = map[string]int{
	"1":  0,
	"b2": 1,
	"2":  2,
	"b3": 3,
	"3":  4,
	"4":  5,
	"b5": 6,
	"5":  7,
	"b6": 8,
	"6":  9,
	"b7": 10,
	"7":  11,
	"b9": 13,
	"9":  14,
}

var patterns = map[string][]string{
	"aeolian":          {"1", "2", "b3", "4", "5", "b6", "b7"},
	"dorian":           {"1", "2", "b3", "4", "5", "6", "b7"},
	"ionian":           {"1", "2", "3", "4", "5", "6", "7"},
	"locrian":          {"1", "b2", "b3", "4", "b5", "b6", "b7"},
	"lydian":           {"1", "2", "3", "b5", "5", "6", "7"},
	"major":            {"1", "2", "3", "4", "5", "6", "7"},
	"major-blues":      {"1", "2", "b3", "3", "5", "6"},
	"major-pentatonic": {"1", "2", "3", "5", "6"},
	"minor":            {"1", "2", "b3", "4", "5", "b6", "b7"},
	"minor-blues":      {"1", "b3", "4", "b5", "5", "b7"},
	"minor-pentatonic": {"1", "b3", "4", "5", "b7"},
	"mixolydian":       {"1", "2", "3", "4", "5", "6", "b7"},
	"phrygian":         {"1", "b2", "b3", "4", "5", "b6", "b7"},

	"5th":         {"1", "5"},
	"aug":         {"1", "3", "b6"},
	"dim":         {"1", "b3", "b5"},
	"m7b5":        {"1", "b3", "b5", "b7"},
	"major-6th":   {"1", "3", "5", "6"},
	"major-7th":   {"1", "3", "5", "b7"},
	"major-maj7":  {"1", "3", "5", "7"},
	"major-triad": {"1", "3", "5"},
	"minor-6th":   {"1", "b3", "5", "6"},
	"minor-7th":   {"1", "b3", "5", "b7"},
	"minor-maj7":  {"1", "b3", "5", "7"},
	"minor-triad": {"1", "b3", "5"},
	"sus":         {"1", "4", "5"},
	"sus2":        {"1", "2", "5"},
	"sus4":        {"1", "4", "5"},
}

type Interval struct {
	name      string
	semitones int
}

func (i Interval) Name() string {
	return i.name
}

func (i Interval) Semitones() int {
	return i.semitones
}

func NewInterval(i string) (*Interval, error) {
	i = normalizeInterval(i)
	_, ok := semitones[i]
	if !ok {
		return nil, errors.New("invalid interval")
	}
	return &Interval{i, semitones[i]}, nil
}

func validInterval(i string) bool {
	i = normalizeInterval(i)
	_, ok := semitones[i]
	return ok
}

func normalizeInterval(i string) string {
	switch i {
	case "P1":
		return "1"
	case "m2":
		return "b2"
	case "M2", "sus2":
		return "2"
	case "m3":
		return "b3"
	case "M3":
		return "3"
	case "P4", "sus4", "sus":
		return "4"
	case "d5":
		return "b5"
	case "P5":
		return "5"
	case "m6", "#5":
		return "b6"
	case "M6":
		return "6"
	case "m7":
		return "b7"
	case "M7":
		return "7"
	case "m9":
		return "b9"
	case "M9":
		return "9"
	default:
		return i
	}
}
