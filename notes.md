# Raft implementation

- Raft is a consensus algorithm

What happens in a Raft network consisting of 0 nodes? What about 1 node?

Test various different kinds of scenarios.

What is a "node"? It's a web server and some internal state.

We just need a program that can act as a node. So basically, that's what we're writing. And it's an executable.

Let's just say that a node has a TCP server in front of it.

---

So I'd like to create a little game.

Interfacet bliver ikke noget særligt lige nu. Det skal bare være nodes, der logger beskeder. Vi kan evt. give log-beskederne en farve hver, så man kan differentiere lidt i, hvad der sker.

Hver node har noget intern state.

- Term count (hvor mange terms har der været?)
- Er jeg leader node?
- Timeout

State machine internal to each node: maybe each node's belief of what the current state is is what we need to store?

Let's just start by implementing one node.

Okay, so

- Find ud af, hvorfor vi ikke kan importere pakken. Det er underligt.
