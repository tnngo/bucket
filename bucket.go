package bucket

import "time"

type token int

const (
	NUMBER token = iota

	// UUID32 32位uuid值, 如果需要追踪请求, 可采用uuid
	UUID32
)

type Bucket struct {
	// count 生产令牌的个数
	count int

	// timer 多少时间内生产令牌
	timer time.Duration

	t token

	// startCode start标记
	// 一旦Bucket.Start(), SetTimer和SetToken不允许再次调用
	startCode int
}

// New 创建Bucket指针对象, 默认每1秒产生count个数字令牌
func New(count int) *Bucket {
	return &Bucket{
		count: count,
		timer: time.Second * 1,
	}
}

// SetTimer 设置多少时间内产生令牌
func (b *Bucket) SetTimer(timer time.Duration) *Bucket {
	if b.startCode == 1 {
		panic("Start后不允许调用SetTimer")
	}
	b.timer = timer
	return b
}

func (b *Bucket) SetToken(t token) *Bucket {
	if b.startCode == 1 {
		panic("Start后不允许调用SetToken")
	}
	b.t = t
	return b
}

// Start 开始放入令牌
func (b *Bucket) Start() {
	b.startCode = 1
}

// Acquire 获得令牌
func (b *Bucket) Acquire() bool {
	return nil
}
