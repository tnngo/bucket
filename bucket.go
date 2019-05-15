package bucket

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tnngo/bucket/queue"
)

type token int

const (
	// NUMBER 默认为数字类型
	NUMBER token = iota

	// UUID32 32位uuid值, 如果需要追踪请求, 可采用uuid
	UUID32
)

type Bucket struct {
	count int
	q     *queue.Queue

	t token
}

// New 创建Bucket指针对象,
// 默认每1秒产生count个令牌
func New(count int, t token) *Bucket {
	if count < 1 {
		panic("每秒生产的令牌数不能小于1")
	}
	return &Bucket{
		count: count,
		q:     queue.New(),
		t:     t,
	}
}

// Start
func (b *Bucket) Start() {
	switch b.t {
	case NUMBER:
		b.addNumber()
	case UUID32:
		b.addUUID32()
	default:
		panic("不支持的token类型")
	}
}

func (b *Bucket) addNumber() {
	n := 0
	for {
		n += 1
		b.q.Put(n)
		time.Sleep(1 * time.Second)
	}
}

func (b *Bucket) addUUID32() {
	uuid32 := uuid.New()
	for {
		b.q.Put(strings.Replace(uuid32.String(), "-", "", -1))
		time.Sleep(1 * time.Second)
	}
}

// Acquire 获得令牌
func (b *Bucket) Acquire() bool {
	return false
}

func (b *Bucket) AcquireNumber() int {
	return 0
}

type Options struct {
	// Count 生产令牌的个数
	Count int

	// Token 令牌类型
	Token token

	// MaxTokenCount 最大令牌数量
	MaxTokenCount int
}

func Custom(options *Options) *Bucket {
	return nil
}
