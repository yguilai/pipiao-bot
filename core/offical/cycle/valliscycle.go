package cycle

import "time"

const (
	vallisCycleDuration = 1600000
	vallisWarmDuration  = 400000
	vallisColdDuration  = vallisCycleDuration - vallisWarmDuration
)

var vallisStartTime, _ = time.Parse("Jan 06, 2006 15:04:05 MST", "November 10, 2018 08:13:48 UTC")

func GetCurrentVallisCycle() *Cycle {
	now := time.Now().UnixMilli()
	sinceLast := (now - vallisStartTime.UnixMilli()) % vallisCycleDuration
	toNextFull := vallisCycleDuration - sinceLast

	state := StateCold
	var millisLeft, prevSub int64
	if toNextFull > vallisColdDuration {
		state = StateWarm
		millisLeft = toNextFull - vallisColdDuration
		prevSub = toNextFull - vallisCycleDuration
	} else {
		millisLeft = toNextFull
		prevSub = toNextFull - vallisColdDuration
	}

	expiry := generateExpiry(now, millisLeft)

	return &Cycle{
		Activation: now + prevSub,
		Expiry:     expiry,
		State:      state,
		TimeLeft:   generateTimeLeft(millisLeft / 1000),
	}
}
