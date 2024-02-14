package main

import (
	"fmt"
	"github.com/maximegosselin/jimi/music"
	"strings"
	"testing"
)

func TestFretboardWithDefaults(t *testing.T) {
	fb, _ := app([]string{})

	want, _ := music.NewNote("E")
	got := fb.GtrStrings()[0].Note()
	if !want.Equals(got) {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func TestChromaticOutputWithDefaults(t *testing.T) {
	got := output([]string{})
	want := []string{
		"E  ┃─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼",
		"B  ┃─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼",
		"G  ┃─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼",
		"D  ┃─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼",
		"A  ┃─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼",
		"E  ┃─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼",
	}
	testOutput(want, got, t)
}

func TestChromaticOutputWithTuning(t *testing.T) {
	got := output([]string{"-t", "DADGbAD"})
	want := []string{
		"D  ┃─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼",
		"A  ┃─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼",
		"F# ┃─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼",
		"D  ┃─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼",
		"A  ┃─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼",
		"D  ┃─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼",
	}
	testOutput(want, got, t)
}

func TestChromaticOutputWithCapo(t *testing.T) {
	got := output([]string{"-c", "4"})
	want := []string{
		"G# ┃────┼────┼────┼──▉─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼",
		"D# ┃────┼────┼────┼──▉─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼",
		"B  ┃────┼────┼────┼──▉─┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼",
		"F# ┃────┼────┼────┼──▉─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼",
		"C# ┃────┼────┼────┼──▉─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼",
		"G# ┃────┼────┼────┼──▉─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼─F──┼─F#─┼─G──┼─G#─┼─A──┼─A#─┼─B──┼─C──┼─C#─┼─D──┼─D#─┼─E──┼",
	}
	testOutput(want, got, t)
}

func TestPresetScaleOutputWithTuningAndCapo(t *testing.T) {
	got := output([]string{"-p", "minor", "-r", "A", "-t", "DADF#AD", "-c", "4"})
	want := []string{
		"F# ┃────┼────┼────┼──▉─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼",
		"C# ┃────┼────┼────┼──▉─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼",
		"A# ┃────┼────┼────┼──▉─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼",
		"F# ┃────┼────┼────┼──▉─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼",
		"C# ┃────┼────┼────┼──▉─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼",
		"F# ┃────┼────┼────┼──▉─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼─D#─┼─E ─┼─F ─┼─F#─┼─G ─┼─G#─┼─A ─┼─A#─┼─B ─┼─C ─┼─C#─┼─D ─┼",
	}
	testOutput(want, got, t)
}

func testOutput(want []string, got []string, t *testing.T) {
	for k, v := range got {
		if want[k] != v {
			t.Fatalf("\nwanted %q\ngot    %q", want[k], v)
		}
	}
}

func output(args []string) []string {
	fb, _ := app(args)
	output := sanitize(fb.String())
	lines := strings.Split(output, "\n")
	return lines[1:7]
}

func sanitize(s string) string {
	good := map[string]struct{}{"A": {}, "B": {}, "C": {}, "D": {}, "E": {}, "F": {}, "G": {}, "#": {}, " ": {}, "┃": {}, "─": {}, "┼": {}, "▉": {}, "\n": {}}
	result := ""
	for _, r := range s {
		char := fmt.Sprintf("%c", r)
		if _, ok := good[char]; ok {
			result += char
		}
	}
	return result
}
