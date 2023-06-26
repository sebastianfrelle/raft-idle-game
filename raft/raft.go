package raft

type Node struct {
	label int
	state int
}

func (node *Node) Broadcast(targets []*Node) {
	for _, n := range targets {
		n.state += 1
	}
}

var n int = 0

type Network []*Node

func NewNetwork(nNodes int) *Network {
	network := make(Network, nNodes)
	for i := range network {
		network[i] = &Node{label: i}
	}

	return &network
}
