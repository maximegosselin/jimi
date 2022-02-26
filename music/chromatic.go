package music

import "github.com/maximegosselin/jimi/ansi"

type ChromaticScale map[string]*Note

func newChromaticScale() ChromaticScale {
	b := &Note{"B", ansi.LightMagenta, nil}
	as := &Note{"A#", ansi.Cyan, b}
	a := &Note{"A", ansi.LightYellow, as}
	gs := &Note{"G#", ansi.Magenta, a}
	g := &Note{"G", ansi.LightGreen, gs}
	fs := &Note{"F#", ansi.Yellow, g}
	f := &Note{"F", ansi.LightRed, fs}
	e := &Note{"E", ansi.Green, f}
	ds := &Note{"D#", ansi.LightBlue, e}
	d := &Note{"D", ansi.Red, ds}
	cs := &Note{"C#", ansi.LightCyan, d}
	c := &Note{"C", ansi.Blue, cs}
	b.next = c
	return ChromaticScale{"C": c, "C#": cs, "D": d, "D#": ds, "E": e, "F": f, "F#": fs, "G": g, "G#": gs, "A": a, "A#": as, "B": b}
}
