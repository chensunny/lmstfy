package engine

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue("test_queue")
	if q.Name() != "test_queue" || q.Tag() != "" {
		t.Fatal("mismatch new queue")
	}

	q = NewQueue("test_queue.{high}")
	if q.Name() != "test_queue" || q.Tag() != "high" {
		t.Fatal("mismatch new queue", q)
	}

	q = NewQueue("t.{low}")
	if q.Name() != "t" || q.Tag() != "low" {
		t.Fatal("mismatch new queue", q)
	}
}
