package main

import (
	"fmt"

	"github.com/nmcapule/quadtree"
	"github.com/nmcapule/yammo/protos/v1/foo"
)

func main() {
	var baz foo.Person
	fmt.Println("hello", baz)

	qt := quadtree.NewQuadtree(quadtree.Bounds{
		X: 0, Y: 0, Width: 100, Height: 100,
	})
	fmt.Println(qt)
}
