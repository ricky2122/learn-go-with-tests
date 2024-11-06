package poker

import (
	"fmt"
	"testing"
	"time"
)

var DummyBlindAlerter = &SpyBlindAlerter{}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

func AssertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

func CheckSchedulingCases(scheduledAlerts []ScheduledAlert, t testing.TB, blindAlerter *SpyBlindAlerter) {
	for i, want := range scheduledAlerts {
		if len(blindAlerter.Alerts) <= i {
			t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
		}

		got := blindAlerter.Alerts[i]
		AssertScheduledAlert(t, got, want)
	}
}
