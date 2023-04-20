package offset

import (
	"github.com/duke-git/lancet/v2/slice"
	"math/rand"
	"sync"
	"time"
)

type OffsetQuota int

const (
	QuotaSubtract OffsetQuota = iota
	QuotaPlus
)

var (
	once   = sync.Once{}
	r      *rand.Rand
	quotas = []OffsetQuota{QuotaSubtract, QuotaPlus}
)

const defaultOffset = 1000

func init() {
	once.Do(func() {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
}

func OffsetInt64(i int64) int64 {
	shuffle := slice.Shuffle(quotas)
	n := r.Int63n(defaultOffset)
	switch shuffle[0] {
	case QuotaSubtract:
		return i - n
	case QuotaPlus:
		return i + n
	}
	return i
}
