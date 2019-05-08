package queue

import (
	"sync"
)

// element 元素
type element struct {
	next  *element
	value interface{}
}

// queue 为令牌桶算法专有队列, 因此设计为有界队列
type queue struct {
	length, total int

	e *element

	mutex *sync.Mutex
}

// New 初始化队列
// total 总数
func New(total int) *queue {
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
	q.e.addElement(v)
	q.length++
}

func (q *queue) take() interface{} {
	q.e = q.e.next
	q.length--
	return nil
}

// Len 返回队列的长度
func (q *queue) Len() int {
	return q.length
}
