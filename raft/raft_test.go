package raft

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	const n int = 3
	network := *NewNetwork(n)

	if len(network) != n {
		t.Fatalf("Expected %d nodes in network; got %d", n, len(network))
	}

	labels := make([]int, len(network))
	for i, n := range network {
		labels[i] = n.label
	}

	exp := []int{0, 1, 2}
	for i := 0; i < n; i++ {
		if labels[i] != exp[i] {
			t.Fatalf("Expected label of node at index %d to be %d; got %d", i, exp[i], labels[i])
		}
	}
}

func TestBroadcast(t *testing.T) {
	network := *NewNetwork(3)
	sender := network[0]
	receivers := network[1:]

	// Assert that the state of every receiver is equal to 0
	for _, r := range receivers {
		if r.state != 0 {
			t.Fatalf("Expected 0; got %d", r.state)
		}
	}

	sender.Broadcast(receivers)

	// Assert that the state of every receiver is equal to 1
	for _, r := range receivers {
		if r.state != 1 {
			t.Fatalf("Expected 1; got %d", r.state)
		}
	}
}
