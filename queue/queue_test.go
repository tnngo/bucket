package queue

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want *queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_element_addElement(t *testing.T) {
	type fields struct {
		next  *element
		value interface{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &element{
				next:  tt.fields.next,
				value: tt.fields.value,
			}
			e.addElement(tt.args.v)
		})
	}
}

func Test_queue_Put(t *testing.T) {
	t.Run("queue.Put", func(t *testing.T) {
		q := New(100)
		q.Put("111")
		q.Put("222")
		q.Put("333")
		t.Logf("%+v", q.e)
		t.Logf("%+v", q.e.next)
		t.Logf("%+v", q.e.next.next)
	})
}

func Test_queue_Len(t *testing.T) {
	type fields struct {
		length     int
		currentLen int
		enqCursor  int
		deqCursor  int
		e          *element
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				length:     tt.fields.length,
				currentLen: tt.fields.currentLen,
				enqCursor:  tt.fields.enqCursor,
				deqCursor:  tt.fields.deqCursor,
				e:          tt.fields.e,
			}
			if got := q.Len(); got != tt.want {
				t.Errorf("queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
