package raindrops

import "strconv"

const (
	testVersion = 3
	dropnoise1  = "Pling"
	dropnoise2  = "Plang"
	dropnoise3  = "Plong"
)

func Convert(dropnumber int) (dropnoise string) {
	if dropnumber%3 == 0 {
		dropnoise += dropnoise1
	}
	if dropnumber%5 == 0 {
		dropnoise += dropnoise2
	}
	if dropnumber%7 == 0 {
		dropnoise += dropnoise3
	}
	if len(dropnoise) == 0 {
		dropnoise = strconv.Itoa(dropnumber)
	}
	return
}
