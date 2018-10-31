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
