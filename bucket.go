package bucket

import (
	"math"
	"time"

	"github.com/tnngo/bucket/queue"
)

type Bucket struct {
	// 最大极限值1000000, 即1微秒1个
	count int

	q *queue.Queue

	duration time.Duration
}

func (b *Bucket) setDuration() *Bucket {
	switch {
	case b.count == 1:
		b.duration = 1 * time.Second
		return b
	case b.count > 1 && b.count <= 1000:
		t := (1.0 / float64(b.count)) * 1000.0
		v := int(math.Floor(t + 0.5))
		b.duration = v * time.Millisecond
		return b
	case b.count > 1000 && b.count <= 1000000:
		t := (1.0 / float64(b.count)) * 1000.0 * 1000.0
		v := int(math.Floor(t + 0.5))
		b.duration = v * time.Microsecond
		return b
	default:
		return b
	}
}

// New 创建Bucket指针对象,
func New(count int) *Bucket {
	switch {
	case count < 1:
		panic("生产的令牌数不能小于1")
	case count > 1000000:
		panic("生产的令牌数不能大于1000000")
	default:
		return (&Bucket{
			count: count,
			q:     queue.New(),
		}).setDuration()
	}
}

// Start
func (b *Bucket) Start() {
	for {
		time.Sleep(b.duration)
	}
}

// Acquire 获得令牌
func (b *Bucket) Acquire() bool {
	return false
}
