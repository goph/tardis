package tardis

import "time"

// Clock provides various time keeping features, such as:
// 		- telling the current time
// 		- sleeping for a given duration
type Clock interface {
	Timepiece
	Sleeper
	TickerFactory
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

// TickerFactory creates a Ticker with a period defined by the duration.
type TickerFactory interface {
	// Ticker creates a Ticker with a period defined by the duration.
	Ticker(d time.Duration) Ticker
}

// Ticker sends a signal (current time) to a channel periodically.
type Ticker interface {
	// Chan returns the signal channel which receives the current time.
	Chan() <-chan time.Time

	// Stop stops the ticker.
	Stop()
}
