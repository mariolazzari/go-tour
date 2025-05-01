package main

import "fmt"

func main() {
	fmt.Println("1 + 2 =", sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}
