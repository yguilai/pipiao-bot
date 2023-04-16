package cycle

import (
	"time"
)

const (
	cetusDayDuration   = 6000000
	cetusNightDuration = 3000000
	cetusNightTime     = 3000
)

func GetCurrentCetusCycle(bounties int64) *Cycle {
	now := time.Now().UnixMilli()
	millisLeft := bounties - now
	secondsToNightEnd := millisLeft / 1000

	var (
		state     CycleState
		subMillis int64
	)
	if secondsToNightEnd > cetusNightTime {
		state = StateDay
		millisLeft = (secondsToNightEnd - cetusNightTime) * 1000
		subMillis = cetusDayDuration
	} else {
		millisLeft = secondsToNightEnd * 1000
		subMillis = cetusNightDuration
	}
	expiry := generateExpiry(now, millisLeft)

	return &Cycle{
		Activation: expiry - subMillis,
		Expiry:     expiry,
		State:      state,
		TimeLeft:   generateTimeLeft(millisLeft / 1000),
	}
}
