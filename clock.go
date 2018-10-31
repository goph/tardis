package tardis

import "time"

// SystemClock uses the real time for telling the time and sleeping.
var SystemClock = ClockFunc(time.Now, time.Sleep)

// Clock tells the current time.
type Clock interface {
	// Now tells the current time.
	Now() time.Time
}

// Sleeper sleeps for a certain amount of time.
type Sleeper interface {
	// Sleep sleeps for a certain amount of time.
	Sleep(d time.Duration)
}

type clockFunc struct {
	nowFn   func() time.Time
	sleepFn func(time.Duration)
}

// ClockFunc returns a new struct from the passed functions.
func ClockFunc(nowFn func() time.Time, sleepFn func(time.Duration)) *clockFunc {
	return &clockFunc{nowFn, sleepFn}
}

// Now tells the current time.
func (c *clockFunc) Now() time.Time {
	return c.nowFn()
}

// Sleep sleeps for a certain amount of time.
func (c *clockFunc) Sleep(d time.Duration) {
	c.sleepFn(d)
}
