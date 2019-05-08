package queue

// element 元素
type element struct {
	next  *element
	value interface{}
}

// queue 为令牌桶算法专有队列, 因此设计为有界队列
type queue struct {
	length, currentLen int

	// 入队游标
	enqCursor int

	// 出队游标
	deqCursor int

	e *element
}

// New 初始化队列
func New(length int) *queue {
	if length < 1 {
		panic("有界队列的长度不可以小于1")
	}
	return &queue{
		length:     length,
		currentLen: 0,
		enqCursor:  0,
		deqCursor:  0,
	}
}

// addElement 添加元素
func (e *element) addElement(v interface{}) {
	for ; e != nil; e = e.next {
		if e.next == nil {
			e.next = &element{
				next:  nil,
				value: v,
			}
			break
		}
	}
}

func (q *queue) Put(v interface{}) {
	if q.Len() == 0 {
		q.e = &element{
			next:  nil,
			value: v,
		}
		q.currentLen = 1
		return
	}
	q.e.addElement(v)
	q.enqCursor++
	q.currentLen++
}

// Len 返回队列的长度
func (q *queue) Len() int {
	return q.currentLen
}
