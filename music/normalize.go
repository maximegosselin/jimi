package music

func normalizeName(n string) string {
	switch n {
	case "B#":
		return "C"
	case "Db":
		return "C#"
	case "Eb":
		return "D#"
	case "Fb":
		return "E"
	case "E#":
		return "F"
	case "Gb":
		return "F#"
	case "Ab":
		return "G#"
	case "Bb":
		return "A#"
	default:
		return n
	}
}
