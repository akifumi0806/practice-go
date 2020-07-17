package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
	fmt.Println(v)
}
