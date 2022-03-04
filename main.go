package main

import "fmt"

const (
	X = iota
	Y
	Z
)

func main() {
	f := integers()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
