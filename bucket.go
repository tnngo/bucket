package bucket

import (
	"fmt"
	"math"
	"time"

	"github.com/tnngo/bucket/queue"
)

type Bucket struct {
	count int

	q *queue.Queue

	duration time.Duration

	baseTicker *time.Ticker
}

func (b *Bucket) setDuration() *Bucket {
	switch {
	case b.count == 1:
		b.duration = 1 * time.Second
		return b
	case b.count > 1 && b.count <= 1000:
		t := (1.0 / float64(b.count)) * 1000.0
		v := time.Duration(math.Floor(t + 0.5))
		b.duration = v * time.Millisecond
		return b
	case b.count > 1000 && b.count <= 1000000:
		t := (1.0 / float64(b.count)) * 1000.0 * 1000.0
		v := time.Duration(math.Floor(t + 0.5))
		b.duration = v * time.Microsecond
		return b
	default:
		return b
	}
}

// start
func (b *Bucket) start() {
	n := 0
	for {
		n++
		b.q.Put(n)
		time.Sleep(b.duration)
	}
}

// New 创建Bucket指针对象
// count 最大极限值1000000, 即1微秒1个.
// 事实上由于硬件或其他一些开销, 比如锁等,
// 并不一定会在1微秒内生成1个, 可能会出现几十或几百毫秒生成1个
func New(count int) *Bucket {
	switch {
	case count < 1:
		panic("生产的令牌数不能小于1")
	case count > 1000000:
		panic("1秒内生成的令牌数不能大于1000000")
	default:
		b := (&Bucket{
			count: count,
			q:     queue.New(),
		}).setDuration()
		go b.start()
		return b
	}
}

func (b *Bucket) Acquire() {
	b.q.Take()
}

/**
 ** 用于限制整个系统流量, 可以用于入口处,
 ** 无论是合法请求还是非法请求,
 ** 只要1个IP在1秒内拿走count/2个令牌,
 ** 则后续其他请求都将进行惩罚用来平衡系统开销,
 ** 直到屏蔽该IP后对其他IP进行速率回复
**/
// EntryAcquire 入口获得令牌
func (b *Bucket) EntryAcquire(ip string) {
	if b.baseTicker == nil {
		b.baseTicker = time.NewTicker(1 * time.Second)
		go func() {
			for {
				select {
				case <-b.baseTicker.C:
					fmt.Println(123)
				}

			}
		}()
	}
	fmt.Println(b.q.Take())
}
