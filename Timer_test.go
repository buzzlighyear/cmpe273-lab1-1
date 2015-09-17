package main

import (
	"testing"
	"time"
)

func PrintSleeptimer(a int) int {
	i := time.Now()
	<-time.After(time.Second * time.Duration(a))
	w := time.Since(i)
	return int(w.Seconds())
}

type testpair struct {
	input  int
	output int
}

var expected = []testpair{
	{1, 1},
	{2, 2},
	{3, 3},
        {4, 4},
        {5, 5},
        {6 ,6},
        {7, 7},
        {8, 8},
        {9, 9},
        {10,10},
}

func TestSleeptimer(t *testing.T) {
	for _, value := range expected {
		v := PrintSleeptimer(value.input)
		if v != value.output {
			t.Errorf("Thread slept for %d seconds instead of  %d", value.input, value.output)
		}
	}
}