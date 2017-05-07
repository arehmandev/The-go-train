package clock

import "fmt"

const testVersion = 4

type Clock int

func New(hour, minute int) Clock {
	time := (hour*60 + minute) % (24 * 60) // convert time to minutes
	if time < 0 {
		time += 60 * 24 // remove negatives
	}
	return Clock(time)
}

func (myclock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", myclock/60, myclock%60)
}

func (myclock Clock) Add(minutes int) Clock {
	newtime := int(myclock) + minutes
	return New(0, newtime)
}
