package guitar

import (
	"fmt"
	"github.com/maximegosselin/jimi/ansi"
	"github.com/maximegosselin/jimi/music"
)

const (
	colorCapo           = ansi.LightWhite
	colorFretBar        = ansi.LightBlack
	colorFretMarkers    = ansi.LightWhite
	colorFretPositions  = ansi.LightWhite
	colorFretString     = ansi.LightBlack
	colorNeckSaddle     = ansi.LightBlack
	colorNoteInScale    = ansi.LightWhite + ansi.BgLightBlack
	colorNoteNotInScale = ansi.LightBlack
	colorScaleKey       = ansi.LightWhite + ansi.BgBlue
	colorStringNote     = ansi.LightWhite
	colorReset          = ansi.Reset
)

type Fretboard struct {
	tuning  Tuning
	strings [6]GtrString
}

func (f Fretboard) Tuning() Tuning {
	return f.tuning
}

func (f Fretboard) GtrStrings() [6]GtrString {
	return f.strings
}

func (f Fretboard) String() string {
	output := colorFretPositions + "     1    2    3    4    5    6    7    8    9    10   11   12   13   14   15   16   17   18   19   20   21   22   23   24\n"
	for i := 0; i <= 5; i++ {
		output += fmt.Sprintln(f.strings[i])
	}
	output += colorFretMarkers + "               ○         ○         ○         ○              ◎              ○         ○         ○         ○              ◎\n" + colorReset
	return output
}

func NewFretboard(t Tuning, s *music.Scale, c uint) Fretboard {
	return Fretboard{
		t,
		[6]GtrString{
			{t[5], s, c},
			{t[4], s, c},
			{t[3], s, c},
			{t[2], s, c},
			{t[1], s, c},
			{t[0], s, c},
		},
	}
}
