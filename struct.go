package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	V := Vertex {1, 2}
	fmt.Println(V.X, V.Y)
	V.X = 4
	fmt.Println(V.X, V.Y, V)
}