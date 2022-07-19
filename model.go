package main

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}
