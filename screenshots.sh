#!/bin/bash

termshot -f docs/defaults.png "./jimi"
termshot -f docs/complete.png "./jimi -t=DADF#AD -k=G -p=triad -c=5"
termshot -f docs/tuning.png "./jimi -t=EbAbDbGbBbEb"
termshot -f docs/capo.png "./jimi -c=4"
termshot -f docs/scale-preset.png "./jimi -k=G# -p=minor-pentatonic"
termshot -f docs/chord-preset.png "./jimi -k=A -p=triad"
termshot -f docs/intervals.png "./jimi -k=C -p=2-2-1-2-2-2-1"
