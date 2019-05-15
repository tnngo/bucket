package bucket

type token int

const (
	NUMBER token = iota

	// UUID32 32位uuid值, 如果需要追踪请求, 可采用uuid
	UUID32
)

type Bucket struct {
	count int
}

// New 创建Bucket指针对象,
// 默认每1秒产生count个数字令牌
func New(count int) *Bucket {
	if count < 1 {
		panic("每秒生产的令牌数不能小于1")
	}
	return &Bucket{
		count: count,
	}
}

type Options struct {
	// Count 生产令牌的个数
	Count int

	// Token 令牌类型
	Token token

	// MaxTokenCount 最大令牌数
	MaxTokenCount int
}

func Custom(options *Options) *Bucket {

}

// Acquire 获得令牌
func (b *Bucket) Acquire() bool {
	return nil
}
