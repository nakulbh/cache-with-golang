package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node //pointer declaration
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}

}

func NewQueue() {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}

}

func main() {
	fmt.Println("Start cache")
	cache := NewCache()
	for _, word := range []string{"parrot", "tree", "patato", "apple", "dog", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}
