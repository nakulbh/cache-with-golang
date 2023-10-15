package main

import "fmt"

type Node struct {
}

type Queue struct {
}

type Hash map[string]*Node

func main() {
	fmt.Println("Start cache")
	cache := Newcache()
	for _, word := range []string{"parrot", "tree", "patato", "apple", "dog", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}

func Newcache() {

}
