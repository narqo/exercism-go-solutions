package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	m := minute % 60
	extraH := minute / 60
	if m < 0 {
		m = 60 + m
		extraH -= 1
	}
	hour += extraH
	h := hour % 24
	if h < 0 {
		h = 24 + h
	}
	return Clock{h, m}
}

func (c Clock) String() string {
	var buf = [4]int{0,0,0,0}
	fmtN(c.h, buf[:2])
	fmtN(c.m, buf[2:])
	return fmt.Sprintf("%d%d:%d%d", buf[0], buf[1], buf[2], buf[3])
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m + minutes)
}

func fmtN(n int, buf []int) {
	if n < 10 {
		buf[1] = n
	} else {
		buf[0] = n / 10
		buf[1] = n % 10
	}
}
