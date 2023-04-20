package cycle

import "time"

const (
	corpusMillis        = 1655182800000
	zarimanDuration     = 18000000
	zarimanDurationHalf = 9000000
)

func GetCurrentZarimanCycle(bounties int64) *Cycle {
	now := time.Now().UnixMilli()
	bounties = bounties - 5000
	millisLeft := bounties - now
	cycleTimeElapsed := (((bounties - corpusMillis) % zarimanDuration) + zarimanDuration) % zarimanDuration
	cycleTimeLeft := zarimanDuration - cycleTimeElapsed
	var state CycleState
	if cycleTimeLeft > zarimanDurationHalf {
		state = StateCorpus
	} else {
		state = StateGrineer
	}
	expiry := generateExpiry(now, millisLeft)
	return &Cycle{
		Activation: expiry - zarimanDurationHalf,
		Expiry:     expiry,
		State:      state,
		TimeLeft:   generateTimeLeft(millisLeft / 1000),
	}
}
