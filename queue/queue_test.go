package queue

import (
	"fmt"
	"testing"
	"time"
)

var q *Queue

func init() {
	q = New(2)
}

func Test_queue_Put(t *testing.T) {
	t.Run("Put3个数据并打印", func(t *testing.T) {
		q := New(2)
		q.Put(111)
		q.Put(222)
		q.Put(333)
		t.Log("结果1: ", q.e.value)
		t.Log("结果2: ", q.e.next.value)
		t.Log("结果3: ", q.e.next.next.value)
		t.Log("长度: ", q.Len())
	})
}

func Test_queue_Take(t *testing.T) {
	t.Run("Take3个数据并打印", func(t *testing.T) {
		q := New(3)
		q.Put(111)
		q.Put(222)
		q.Put(333)
		t.Log("结果1: ", q.Take())
		t.Log("结果2: ", q.Take())
		t.Log("结果3: ", q.Take())
	})
}

// Test_queue_Take_for_blocking 测试阻塞操作
func Test_queue_Take_for_blocking(t *testing.T) {
	q := New(6)
	go func() {
		t.Run("Put", func(t *testing.T) {
			for {
				q.Put(1)
				q.Put(2)
				q.Put(3)
				q.Put(4)
				q.Put(5)
				q.Put(6)
				time.Sleep(time.Second * 1)
			}
		})
	}()

	go func() {
		t.Run("Take", func(t *testing.T) {
			for {
				fmt.Println(q.Take())
			}
		})
	}()
	select {}
}
