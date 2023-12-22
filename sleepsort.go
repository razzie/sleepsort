package sleepsort

import "time"

func SleepSort(x []uint) {
	c := make(chan uint)
	for _, n := range x {
		n := n
		time.AfterFunc(time.Duration(n), func() { c <- n })
	}
	for i := range x {
		x[i] = <-c
	}
}
