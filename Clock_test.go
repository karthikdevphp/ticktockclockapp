package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestClock(t *testing.T) {

	//wantedStr := "Elapsed: 1 Tick\nElapsed: 2 Tick\nElapsed: 3 Tick\nElapsed: 4 Tick\nElapsed: 5 Tick\nElapsed: 6 Tick\nElapsed: 7 Tick\nElapsed: 8 Tick\nElapsed: 9 Tick\nElapsed: 10 Tick\nElapsed: 10 Tock\nElapsed: 11 Tick\nElapsed: 12 Tick\nElapsed: 13 Tick\nElapsed: 14 Tick\nElapsed: 15 Tick\nElapsed: 16 Tick\nElapsed: 17 Tick\nElapsed: 18 Tick\nElapsed: 19 Tick\nElapsed: 20 Tick\nElapsed: 20 Tock\nElapsed: 21 Tick\nElapsed: 22 Tick\nElapsed: 23 Tick\nElapsed: 24 Tick\nElapsed: 25 Tick\nElapsed: 26 Tick\nElapsed: 27 Tick\nElapsed: 28 Tick\nElapsed: 29 Tick\nElapsed: 30 Tick\nElapsed: 30 Tock\nElapsed: 30 Bong\n"
	wantedStr := "Elapsed: 1 Tick\nElapsed: 2 Tick\nElapsed: 3 Tick\nElapsed: 4 Tick\nElapsed: 5 Tick\nElapsed: 6 Tick\nElapsed: 7 Tick\nElapsed: 8 Tick\nElapsed: 9 Tick\nElapsed: 10 Tick\nElapsed: 10 Tock\nElapsed: 11 Tick\nElapsed: 12 Tick\nElapsed: 13 Tick\nElapsed: 14 Tick\nElapsed: 15 Tick\nElapsed: 16 Tick\nElapsed: 17 Tick\nElapsed: 18 Tick\nElapsed: 19 Tick\nElapsed: 20 Tick\nElapsed: 20 Tock\nElapsed: 21 Tick\nElapsed: 22 Tick\nElapsed: 23 Tick\nElapsed: 24 Tick\nElapsed: 25 Tick\nElapsed: 26 Tick\nElapsed: 27 Tick\nElapsed: 28 Tick\nElapsed: 29 Tick\nElapsed: 30 Tick\nElapsed: 30 Bong\n"
	var tests = []struct {
		sec  time.Duration
		min  time.Duration
		hour time.Duration
		stopTime time.Duration
		want string
		}{
			{1, 10, 30, 30, wantedStr},
		}

	for _, test := range tests {
		descr := fmt.Sprintf("clockStart(%s,%s,%s,%s)", test.sec, test.min, test.hour,test.stopTime)
		out = new(bytes.Buffer)
		if err := clockStart(test.sec,test.min,test.hour,test.stopTime); err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q,want %q", descr, got, test.want)
		}
	}
}



