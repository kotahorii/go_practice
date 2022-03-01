package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// func fibonacchi() func() int {
// 	a, b := 1, 0
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

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
