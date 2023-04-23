package cycle

import (
	"math"
)

const (
	minutesCoef = 6000
)

type CycleState string

const (
	StateDay     CycleState = "day"
	StateNight   CycleState = "night"
	StateCold    CycleState = "cold"
	StateWarm    CycleState = "warm"
	StateCorpus  CycleState = "corpus"
	StateGrineer CycleState = "grineer"
)

type Cycle struct {
	Id         string
	Activation int64
	Expiry     int64
	State      CycleState
	TimeLeft   []int64
}

func generateTimeLeft(secondsLeft int64) []int64 {
	timeLeft := make([]int64, 3)
	if secondsLeft > 3600 {
		timeLeft[0] = secondsLeft / 3600
		secondsLeft %= 3600
	}
	if secondsLeft > 60 {
		timeLeft[1] = secondsLeft / 60
		secondsLeft %= 60
	}
	timeLeft[2] = secondsLeft
	return timeLeft
}

func generateExpiry(now, millisLeft int64) int64 {
	return int64(math.Round(float64((now+millisLeft)/minutesCoef)) * minutesCoef)
}
