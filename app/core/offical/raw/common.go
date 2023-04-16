package raw

import (
	"strconv"
	"time"
)

type IdObj struct {
	Id string `json:"$oid"`
}

type TimeObj struct {
	Date struct {
		NumberLong string `json:"$numberLong"`
	} `json:"$date"`
}

func (t *TimeObj) Time() (*time.Time, error) {
	millis, err := t.Millis()
	if err != nil {
		return nil, err
	}
	ti := time.UnixMilli(millis)
	return &ti, nil
}

func (t *TimeObj) Millis() (int64, error) {
	return strconv.ParseInt(t.Date.NumberLong, 10, 64)
}
