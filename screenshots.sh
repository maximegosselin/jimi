#!/bin/bash

termshot -f docs/defaults.png "./jimi"
termshot -f docs/complete.png "./jimi -t=DADF#AD -r=G -p=major-triad -c=5"
termshot -f docs/tuning.png "./jimi -t=EbAbDbGbBbEb"
termshot -f docs/capo.png "./jimi -c=4"
termshot -f docs/scale-preset.png "./jimi -r=G# -p=minor-pentatonic"
termshot -f docs/chord-preset.png "./jimi -r=A -p=major-triad"
termshot -f docs/intervals.png "./jimi -r=C -p=1-3-5-b7"
