package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 5))
}

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Times(a int, b int) int {
	return a * b
}

func SumX(a int, b int) int {
	return a + b + a
}
