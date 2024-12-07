package stacks

import "testing"

// TestLinkedStack tests the LinkedStack.
func TestLinkedStack(t *testing.T) {
	const (
		MAX int = 1000
	)

	want := make([]int, 0, MAX+1)

	for i := 0; i < MAX+1; i++ {
		want = append(want, i)
	}

	stack := new(LinkedStack[int])

	for _, e := range want {
		err := stack.Push(e)
		if err != nil {
			t.Fatalf("failed to push: %s", err.Error())
		}
	}

	var got []int

	for {
		e, err := stack.Pop()
		if err != nil {
			break
		}

		got = append(got, e)
	}

	if len(got) != len(want) {
		t.Fatalf("got %d, want %d", len(got), len(want))
	}
}
