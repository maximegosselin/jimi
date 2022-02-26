package guitar

import (
	"fmt"
	"github.com/maximegosselin/jimi/music"
	"strings"
)

const (
	stringChar  = "─"
	fretBarChar = "┼"
	capoBarChar = "▉"
)

type GtrString struct {
	note  *music.Note
	scale *music.Scale
	capo  uint
}

func (gs GtrString) Note() *music.Note {
	return gs.note
}

func (gs GtrString) Capo() uint {
	return gs.capo
}

func (gs GtrString) String() string {
	chars := ""
	note := gs.note.Next()
	for i := 0; i < 24; i++ {
		if i < int(gs.capo-1) {
			chars += gs.emptyFret()
		} else if i == int(gs.capo-1) {
			chars += gs.capoFret()
		} else {
			chars += gs.noteFret(note)
		}
		note = note.Next()
	}
	return gs.neck() + chars
}

func (gs GtrString) neck() string {
	note := gs.note
	for i := 0; i < int(gs.capo); i++ {
		note = note.Next()
	}
	return colorStringNote + fmt.Sprintf("%-3s", note.Name()) + colorNeckSaddle + "┃" + colorReset
}

func (gs GtrString) noteFret(note *music.Note) string {
	chars := colorFretString +
		stringChar

	if gs.scale.Name() == "chromatic" {
		chars += note.Color() +
			note.Name() +
			colorReset +
			colorFretString +
			strings.Repeat(stringChar, -len(note.Name())+2) +
			stringChar +
			colorFretBar +
			fretBarChar +
			colorReset
		return chars
	}

	if !gs.scale.ContainsNote(note) {
		chars += colorNoteNotInScale +
			note.Name() +
			colorFretString +
			strings.Repeat(stringChar, -len(note.Name())+2)
	} else if gs.scale.KeyIs(note) {
		chars += colorScaleKey +
			note.Name() +
			strings.Repeat(" ", -len(note.Name())+2)
	} else {
		chars += colorNoteInScale +
			note.Name() +
			strings.Repeat(" ", -len(note.Name())+2)
	}
	chars += colorReset +
		colorFretString +
		stringChar +
		colorFretBar +
		fretBarChar +
		colorReset

	return chars
}

func (gs GtrString) capoFret() string {
	return colorFretString +
		stringChar +
		stringChar +
		colorCapo +
		capoBarChar +
		colorFretString +
		stringChar +
		colorFretBar +
		fretBarChar +
		colorReset
}

func (gs GtrString) emptyFret() string {
	return colorFretString +
		stringChar +
		stringChar +
		stringChar +
		stringChar +
		colorFretBar +
		fretBarChar +
		colorReset
}
