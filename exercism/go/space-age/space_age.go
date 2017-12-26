package space

// Planet - name of planet
type Planet string

// divider - used to calculate age
var divider float64

// Age - converts age from seconds to years on each planet
func Age(seconds float64, planet Planet) float64 {

	switch planet {
	case "Earth":
		divider = 1000000000 / 31.69
	case "Mercury":
		divider = 2134835688 / 280.88
	case "Venus":
		divider = 189839836 / 9.78
	case "Mars":
		divider = 2329871239 / 39.25
	case "Jupiter":
		divider = 901876382 / 2.41
	case "Saturn":
		divider = 3000000000 / 3.23
	case "Uranus":
		divider = 3210123456 / 1.21
	case "Neptune":
		divider = 8210123456 / 1.58
	default:
		return 0
	}

	return seconds / divider
}
