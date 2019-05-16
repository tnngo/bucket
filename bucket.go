package bucket

import (
	"github.com/tnngo/bucket/queue"
)

type Bucket struct {
	count int

	q *queue.Queue
}

// New 创建Bucket指针对象,
func New(count int) *Bucket {
	if count < 1 {
		panic("生产的令牌数不能小于1")
	}
	return &Bucket{
		count: count,
		q:     queue.New(),
	}
}

// Start
func (b *Bucket) Start() {
}

// Acquire 获得令牌
func (b *Bucket) Acquire() bool {
	return false
}
