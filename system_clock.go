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

// Ticker creates a Ticker with a period defined by the duration.
func (*systemClock) Ticker(d time.Duration) Ticker {
	return &ticker{time.NewTicker(d)}
}

type ticker struct {
	t *time.Ticker
}

func (t *ticker) Chan() <-chan time.Time {
	return t.t.C
}

func (t *ticker) Stop() {
	t.t.Stop()
}
