package main

import (
	"flag"
	"github.com/maximegosselin/jimi/guitar"
	"github.com/maximegosselin/jimi/music"
)

func app(args []string) (*guitar.Fretboard, error) {
	fs := flag.NewFlagSet("jimi", flag.ExitOnError)
	capoArg := fs.Uint("c", 0, "Fret position of capo.")
	rootArg := fs.String("r", "C", "Root note.")
	patternArg := fs.String("p", "chromatic", "Pattern name (ie. \"minor\") or intervals (ie. \"1-b3-5\").")
	tuningArg := fs.String("t", "EADGBE", "Guitar tuning from low to high.")
	fs.Parse(args)

	t, err := guitar.NewTuning(*tuningArg)
	if err != nil {
		return nil, err
	}
	r, err := music.NewNote(*rootArg)
	if err != nil {
		return nil, err
	}
	s := music.NewScale(*patternArg, r)

	fb := guitar.NewFretboard(t, s, *capoArg)
	return &fb, nil
}
