package main

import "fmt"

const SIZE = 5

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail, Length: 1}
}

func (c *Cache) Check(str string) {
	var node *Node
	if val := c.Hash[str]; val != nil {
		node = c.Remove(val)
	} else {
		node = &Node{Value: str}
	}

	c.Add(node)

	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove : %s \n", n.Value)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1

	delete(c.Hash, n.Value)

	return n
}

func (c *Cache) Add(n *Node) *Node {
	fmt.Printf("add : %s \n", n.Value)
	temp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Right = temp
	n.Left = c.Queue.Head
	temp.Left = n

	c.Queue.Length++

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

	return n
}

func (c Cache) Display() {
	c.Queue.Display()
}

func (q Queue) Display() {
	node := q.Head.Right

	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {
		fmt.Printf("%s}", node.Value)

		if i < q.Length-1 {
			fmt.Printf("<-->")
		}

		node = node.Right
	}

	fmt.Printf("]\n")
}

func main() {
	fmt.Println("Cache has started")
	cache := NewCache()

	for _, word := range []string{"parrot", "avocaod", "dragonfruit", "tree", "potato", "tomato", "tree", "parrot"} {
		cache.Check(word)
		cache.Display()
	}
}
