package queue

import (
	"sync"
	"sync/atomic"
)

// element 元素
type element struct {
	next *element

	value interface{}
}

// queue 为令牌桶算法专有队列, 因此设计为有界队列
type queue struct {
	length, total int32

	e *element

	mutex *sync.Mutex
}

// New 初始化队列
// total 总数
func New(total int32) *queue {
	if total < 1 {
		panic("有界队列的长度不可以小于1")
	}
	return &queue{
		total:  total,
		length: 0,
		mutex:  &sync.Mutex{},
	}
}

// addElement 添加元素
func (e *element) add(v interface{}) {
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

// remove 删除队头, 并返回next指针地址
func (e *element) remove() (*element, interface{}) {
	next := e.next
	v := e.value
	e = nil
	return next, v
}

// Put 入队, 阻塞式操作
func (q *queue) Put(v interface{}) {
	if q.Len() == 0 {
		q.e = &element{
			next:  nil,
			value: v,
		}
		q.length = 1
		return
	}
	q.e.add(v)
	// TODO 考虑在mutex.Lock内进行+1
	newLen := atomic.AddInt32(&q.length, 1)
	q.length = newLen
}

// Take 出队, 阻塞式操作
func (q *queue) Take() interface{} {
	var v interface{}
	q.e, v = q.e.remove()
	// TODO 考虑在mutex.Lock内进行+1
	newLen := atomic.AddInt32(&q.length, -1)
	q.length = newLen
	return v
}

// Len 返回队列的长度
func (q *queue) Len() int32 {
	return q.length
}
