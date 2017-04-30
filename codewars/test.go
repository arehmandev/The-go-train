package main

import (
	"fmt"
	"sort"
)

func main() {

	var slice []float64
	v := 25
	g := 9.81

	for i := 0; i < v; i++ {
		v0 := float64(v) * 1000 / 60 / 60
		index := float64(i)
		t := index / 10
		h := (v0 * t) - (0.5 * g * t * t)
		slice = append(slice, h)
		sort.Sort(sort.Reverse(sort.Float64Slice(slice)))

		highest := slice[0]

		if h < highest {
			fmt.Println("Distance:", highest, "Time", t)
			break
		}
	}
}

// experimental code
//sort.Float64s(slice)

/*
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
*/

//	fmt.Println(slice)
