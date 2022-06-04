package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z = x
	i := 1
	for i < 10 {
		z -= (z*z - x) / (2*z)
		i ++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}