package music

import (
	"errors"
)

var semitones = map[string]int{
	"R":  0,
	"m2": 1,
	"M2": 2,
	"m3": 3,
	"M3": 4,
	"P4": 5,
	"d5": 6,
	"P5": 7,
	"m6": 8,
	"M6": 9,
	"m7": 10,
	"M7": 11,
}

var patterns = map[string][]string{
	"aeolian":          {"R", "M2", "m3", "P4", "P5", "m6", "m7"},
	"dorian":           {"R", "M2", "m3", "P4", "P5", "M6", "m7"},
	"ionian":           {"R", "M2", "M3", "P4", "P5", "M6", "M7"},
	"locrian":          {"R", "m2", "m3", "P4", "d5", "m6", "m7"},
	"lydian":           {"R", "M2", "M3", "d5", "P5", "M6", "M7"},
	"major":            {"R", "M2", "M3", "P4", "P5", "M6", "M7"},
	"major-blues":      {"R", "M2", "m3", "M3", "P5", "M6"},
	"major-pentatonic": {"R", "M2", "M3", "P5", "M6"},
	"minor":            {"R", "M2", "m3", "P4", "P5", "m6", "m7"},
	"minor-blues":      {"R", "m3", "P4", "d5", "P5", "m7"},
	"minor-pentatonic": {"R", "m3", "P4", "P5", "m7"},
	"mixolydian":       {"R", "M2", "M3", "P4", "P5", "M6", "m7"},
	"phrygian":         {"R", "m2", "m3", "P4", "P5", "m6", "m7"},

	"5th":         {"R", "P5"},
	"aug":         {"R", "M3", "m6"},
	"dim":         {"R", "m3", "d5"},
	"m7b5":        {"R", "m3", "d5", "m7"},
	"major-6th":   {"R", "M3", "P5", "M6"},
	"major-7th":   {"R", "M3", "P5", "m7"},
	"major-maj7":  {"R", "M3", "P5", "M7"},
	"major-triad": {"R", "M3", "P5"},
	"minor-6th":   {"R", "m3", "P5", "M6"},
	"minor-7th":   {"R", "m3", "P5", "m7"},
	"minor-maj7":  {"R", "m3", "P5", "M7"},
	"minor-triad": {"R", "m3", "P5"},
	"sus":         {"R", "P4", "P5"},
	"sus2":        {"R", "M2", "P5"},
	"sus4":        {"R", "P4", "P5"},
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
	case "1", "P1":
		return "R"
	case "b2":
		return "m2"
	case "2", "sus2":
		return "M2"
	case "b3":
		return "m3"
	case "3":
		return "M3"
	case "4", "sus4", "sus":
		return "P4"
	case "b5":
		return "d5"
	case "5":
		return "P5"
	case "b6", "#5":
		return "m6"
	case "6":
		return "M6"
	case "b7":
		return "m7"
	case "7":
		return "M7"
	case "b9", "m9":
		return "m2"
	case "9", "M9":
		return "2"
	default:
		return i
	}
}
