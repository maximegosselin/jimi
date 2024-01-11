package music

func normalizeName(n string) string {
	switch n {
	case "Cb":
		return "B"
	case "Db":
		return "C#"
	case "Eb":
		return "D#"
	case "E#":
		return "F"
	case "Fb":
		return "E"
	case "Gb":
		return "F#"
	case "Ab":
		return "G#"
	case "Bb":
		return "A#"
	case "B#":
		return "C"
	default:
		return n
	}
}
