package main

import (
	"fmt"
	"time"
)

type TimeADM struct {
	LargestTime  time.Duration
	ShortestTime time.Duration
	Times        []time.Duration
}

func (t *TimeADM) Add(interval time.Duration) {
	if interval > t.LargestTime {
		t.LargestTime = interval
	}
	if interval < t.ShortestTime {
		t.ShortestTime = interval
	}
	t.Times = append(t.Times, interval)
}

func (t TimeADM) Show(op string) string {
	return fmt.Sprintf("%s: Shortest: %v; Largest: %v", op, t.ShortestTime, t.LargestTime)
}
