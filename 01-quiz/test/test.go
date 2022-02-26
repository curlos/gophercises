package main

import (
	"fmt"
	"time"
)

func main() {
	sec := 30 + 2
	deadline := time.Now().Add(time.Second * time.Duration(sec))

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(deadline)

		if timeRemaining.t <= 0 {
			fmt.Println("Countdown reached!")
			break
		}

		fmt.Printf("Seconds: %d\n", timeRemaining.s)
	}
}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
