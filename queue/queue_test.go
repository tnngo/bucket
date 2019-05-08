package queue

import (
	"testing"
)

func Test_queue_Put(t *testing.T) {
	t.Run("Put3个数据并打印", func(t *testing.T) {
		q := New(100)
		q.Put(111)
		q.Put(222)
		q.Put(333)
		t.Log("结果1: ", q.e.value)
		t.Log("结果2: ", q.e.next.value)
		t.Log("结果3: ", q.e.next.next.value)
	})
}
