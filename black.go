package bucket

import (
	"time"
)

type Black struct {
	b      *Bucket
	ticker *time.Ticker
}

func NewBlack(count int) *Black {
	return &Black{
		b:      New(count),
		ticker: time.NewTicker(1 * time.Second),
	}
}

// RedisOption Redis配置
type RedisOption struct {
	IP       string
	User     string
	Password string
}

func (black *Black) SetRedis(ro *RedisOption) {}

/**
 ** 用于限制整个系统流量, 可以用于入口处,
 ** 无论是合法请求还是非法请求,
 ** 只要1个IP在1秒内拿走count/2个令牌,
 ** 则后续其他请求都将进行惩罚用来平衡系统开销,
 ** 直到屏蔽该IP后对其他IP进行速率恢复.
**/
// Acquire
func (black *Black) Acquire(ip string) {
	go func() {
		for {
			select {
			case <-black.ticker.C:
			default:
				black.ticker.Stop()
			}

		}
	}()
	black.b.Acquire()
}
