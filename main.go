package main

import "fmt"

func main() {
	fib := fibonacchi()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func fibonacchi() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func printSlice(x []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
// }

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	return lim
// }

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for y := 0.0; y <= 10; y++ {
// 		z -= (z*z - y) / (2 * z)
// 		fmt.Println(z / math.Sqrt(x))
// 	}
// 	return z
// }
