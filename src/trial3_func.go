package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func swap_naked(x, y string) (m, n string) {
	m = y
	n = x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(swap_naked("hello", "world"))
}
