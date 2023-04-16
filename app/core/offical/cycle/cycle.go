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
	StateNight              = "night"
	StateCold               = "cold"
	StateWarm               = "warm"
	StateCorpus             = "corpus"
	StateGrineer            = "grineer"
)

type Cycle struct {
	Id         string
	Activation int64
	Expiry     int64
	State      CycleState
	TimeLeft   []int
}

func generateTimeLeft(secondsLeft int64) []int {
	timeLeft := make([]int, 3)
	if secondsLeft > 3600 {
		timeLeft[0] = int(math.Floor(float64(secondsLeft / 3600)))
		secondsLeft %= 3600
	}
	if secondsLeft > 60 {
		timeLeft[1] = int(math.Floor(float64(secondsLeft / 60)))
		secondsLeft %= 60
	}
	timeLeft[2] = int(secondsLeft)
	return timeLeft
}

func generateExpiry(now, millisLeft int64) int64 {
	return int64(math.Round(float64((now+millisLeft)/minutesCoef)) * minutesCoef)
}
