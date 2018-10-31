package tardis

import "time"

// SystemClock uses the real time.
var SystemClock = NewSystemClock()

type systemClock struct{}

// NewSystemClock returns a clock that uses real time.
func NewSystemClock() Clock {
	return &systemClock{}
}

// Now tells the current time.
func (*systemClock) Now() time.Time {
	return time.Now()
}

// Sleep sleeps for a certain amount of time.
func (*systemClock) Sleep(d time.Duration) {
	time.Sleep(d)
}
