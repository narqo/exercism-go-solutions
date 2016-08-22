package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

const minutesPerDay = 24 * 60

// Clock represents a clock that handles time without date.
type Clock int

// New creates a new clock.
func New(hour, minute int) Clock {
	m := (hour*60 + minute) % minutesPerDay
	if m < 0 {
		m += minutesPerDay
	}
	return Clock(m)
}

// String returns a string representation of the clock.
func (c Clock) String() string {
	h := c / 60
	return fmt.Sprintf("%02d:%02d", h, c-h*60)
}

// Add creates a new clock that is `minute` ahead from this one.
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
