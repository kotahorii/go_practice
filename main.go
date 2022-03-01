package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
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
