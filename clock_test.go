package tardis

import (
	"testing"
	"time"
)

func TestClockFunc_Now(t *testing.T) {
	now, _ := time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 2015")

	clock := ClockFunc(
		func() time.Time {
			return now
		},
		nil,
	)

	if got, want := clock.Now(), now; got != want {
		t.Errorf("expected %q, got: %q", want, got)
	}
}

func TestClockFunc_Sleep(t *testing.T) {
	var duration time.Duration

	clock := ClockFunc(
		nil,
		func(d time.Duration) {
			duration = d
		},
	)

	clock.Sleep(time.Second)

	if got, want := duration, time.Second; got != want {
		t.Errorf("expected %q, got: %q", want, got)
	}
}
