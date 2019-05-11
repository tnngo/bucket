package queue

import (
	"sync"
)

// element 元素
type element struct {
	next *element

	value interface{}
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
	// 本不该对e == nil进行判断, 但神奇的Testing会偶尔破坏互斥性
	// 因此为了安全起见对e == nil 进行判断
	if e == nil {
		return nil, nil
	}
	next, v := e.next, e.value
	e = nil
	return next, v
}

// queue 为令牌桶算法专有队列, 因此设计为有界队列
type Queue struct {
	length, total int32

	e *element

	mutexPut  sync.Mutex
	mutexTake sync.Mutex

	notifyPut  *sync.Cond
	notifyTake *sync.Cond
}

// New 初始化队列
// total 总数
func New(total int32) *Queue {
	if total < 1 {
		panic("有界队列的长度不可以小于1")
	}
	q := &Queue{
		total:  total,
		length: 0,
	}
	q.notifyPut = sync.NewCond(&q.mutexPut)
	q.notifyTake = sync.NewCond(&q.mutexTake)
	return q
}

// Put 入队, 阻塞式操作
func (q *Queue) Put(v interface{}) {
	if q.Len() == 0 {
		q.e = &element{
			next:  nil,
			value: v,
		}
		q.length = 1
		return
	}

	q.mutexPut.Lock()
	defer q.mutexPut.Unlock()
	for q.length == q.total {
		q.notifyPut.Wait()
	}

	q.e.add(v)
	q.length++
	q.notifyTake.Signal()
}

// Take 出队, 阻塞式操作
func (q *Queue) Take() interface{} {
	q.mutexTake.Lock()
	defer q.mutexTake.Unlock()
	if q.length == 0 {
		q.notifyTake.Wait()
	}
	var v interface{}
	q.e, v = q.e.remove()
	q.length--
	q.notifyPut.Signal()
	return v
}

// Len 返回队列的长度
func (q *Queue) Len() int32 {
	return q.length
}
