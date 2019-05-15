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
	next, v := e.next, e.value
	e = nil
	return next, v
}

// queue 为令牌桶算法专有队列
type Queue struct {
	// total 如果tatal等于0, 则为无界队列
	total  int
	length int

	e *element

	mutex sync.Mutex

	notifyPut  *sync.Cond
	notifyTake *sync.Cond
}

// New 初始化队列
// total 总数
func New() *Queue {
	q := &Queue{
		length: 0,
	}

	q.notifyPut = sync.NewCond(&q.mutex)
	q.notifyTake = sync.NewCond(&q.mutex)
	return q
}

// Put 入队, 如果队列满,则进行阻塞
func (q *Queue) Put(v interface{}) {
	if v == nil {
		panic("不能Put空值")
	}
	defer q.mutex.Unlock()
	q.mutex.Lock()
	if q.Len() == 0 {
		q.e = &element{
			next:  nil,
			value: v,
		}
		q.length = 1
		return
	}

	q.e.add(v)
	q.length += 1
	for q.length == q.total {
		q.notifyPut.Wait()
	}

	q.notifyTake.Signal()
}

// Take 出队, 如果队列为空, 则进行阻塞
// 内存不够时, 互斥会失败
func (q *Queue) Take() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for q.length == 0 {
		q.notifyTake.Wait()
	}
	var v interface{}
	q.e, v = q.e.remove()
	q.length -= 1
	q.notifyPut.Signal()
	return v
}

// Len 返回队列的长度
func (q *Queue) Len() int {
	return q.length
}
