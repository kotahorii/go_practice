package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println((compute(math.Pow)))
}

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
