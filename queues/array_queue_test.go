package queues

import (
	"slices"
	"testing"
)

// TestLinkedQueue tests the LinkedQueue.
func TestArrayQueue(t *testing.T) {
	queue := new(ArrayQueue[uint])

	const (
		MAX uint = 1000
	)

	want := make([]uint, 0, MAX+1)

	for i := uint(0); i < MAX+1; i++ {
		want = append(want, i)
	}

	for _, e := range want {
		err := queue.Enqueue(e)
		if err != nil {
			t.Fatalf("failed to enqueue: %s", err.Error())
		}
	}

	var got []uint

	for {
		e, err := queue.Dequeue()
		if err != nil {
			break
		}

		got = append(got, e)
	}

	ok := slices.Equal(want, got)
	if !ok {
		t.Fatalf("want %v, got %v", want, got)
	}
}
