package main

import "fmt"

func main() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	fmt.Println(p1)
	p2 = &j
	fmt.Println(p2)
	i = *p1 + *p2
	fmt.Println(i)
	p2 = p1
	fmt.Println(p2)
	j = *p2 + i
	fmt.Println(j)
}
