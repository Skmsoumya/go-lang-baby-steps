package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += 1
	}
	fmt.Println("sum", sum);
}