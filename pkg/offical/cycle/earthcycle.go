package cycle

import (
	"time"
)

const (
	earthCycleSecond     = 28800
	earthCycleSecondHalf = earthCycleSecond / 2
	fourHourCoef         = 1000 * 60 * 60 * 4
)

func GetCurrentEarthCycle() *Cycle {
	now := time.Now().UnixMilli()
	cycleSeconds := now / 1000 % earthCycleSecond

	secondsLeft := earthCycleSecondHalf - (cycleSeconds % earthCycleSecondHalf)
	millisLeft := secondsLeft * 1000
	expiry := generateExpiry(now, millisLeft)

	state := StateDay
	if cycleSeconds > earthCycleSecondHalf {
		state = StateNight
	}
	return &Cycle{
		Activation: expiry - fourHourCoef,
		Expiry:     expiry,
		State:      state,
		TimeLeft:   generateTimeLeft(secondsLeft),
	}
}
