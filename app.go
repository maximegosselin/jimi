package main

import (
	"flag"
	"github.com/maximegosselin/jimi/guitar"
	"github.com/maximegosselin/jimi/music"
)

func app(args []string) (*guitar.Fretboard, error) {
	fs := flag.NewFlagSet("jimi", flag.ExitOnError)
	capoArg := fs.Uint("c", 0, "Fret position of capo.")
	keyArg := fs.String("k", "C", "Key of the scale.")
	patternArg := fs.String("p", "chromatic", "Pattern name (ie. \"major\") or semitone intervals (ie. \"2-2-1-2-2-2-1\").")
	tuningArg := fs.String("t", "EADGBE", "Guitar tuning from low to high.")
	fs.Parse(args)

	t, err := guitar.NewTuning(*tuningArg)
	if err != nil {
		return nil, err
	}
	k, err := music.NewNote(*keyArg)
	if err != nil {
		return nil, err
	}
	s := music.NewScale(*patternArg, k)

	fb := guitar.NewFretboard(t, s, *capoArg)
	return &fb, nil
}
