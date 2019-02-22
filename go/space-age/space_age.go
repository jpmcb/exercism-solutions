package space

// Planet is the name of the planet
type Planet string

// Age takes the name of a planet and seconds. Returns the earth years
func Age(seconds float64, name Planet) float64 {
	earthYears := seconds / 31557600
	switch name {
	case "Earth":
		return earthYears
	case "Mercury":
		return earthYears / 0.2408467
	case "Venus":
		return earthYears / 0.61519726
	case "Mars":
		return earthYears / 1.8808158
	case "Jupiter":
		return earthYears / 11.862615
	case "Saturn":
		return earthYears / 29.447498
	case "Uranus":
		return earthYears / 84.016846
	case "Neptune":
		return earthYears / 164.79132
	default:
		return 0
	}
}
