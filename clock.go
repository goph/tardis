package tardis

import "time"

// Clock provides various time keeping features, such as:
// 		- telling the current time
// 		- sleeping for a given duration
type Clock interface {
	Timepiece
	Sleeper
}

// Timepiece tells the current time.
type Timepiece interface {
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
func ClockFunc(nowFn func() time.Time, sleepFn func(time.Duration)) Clock {
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
