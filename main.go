package main

import "fmt"

// V 1.0.0
func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(3, 4)
	fmt.Println(result)
}
