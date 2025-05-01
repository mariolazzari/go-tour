package main

import "fmt"

func main() {
	a, b := swap("Mario", "Lazzari")
	fmt.Println(a, b)
}

func swap(a, b string) (string, string) {
	return b, a
}
