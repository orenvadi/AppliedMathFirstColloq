package main

import (
	"fmt"
	"math"
)

func main() {
	arguments := []float64{8.45, 23.6, 1 / 4.09, 0.79, math.Pow(math.E, 2.01), math.Pow(3.4, 2.6)}

	for i, arg := range arguments {
		switch i {
		case 0:
			// arctg 8.45
			result := math.Atan(arg)
			fmt.Printf("arctg(%.4f) = %.4f\n", arg, result)
		case 1:
			// lg 23.6
			result := math.Log10(arg)
			fmt.Printf("lg(%.4f) = %.4f\n", arg, result)
		case 2:
			// 1/4.09
			result := 1 / arg
			fmt.Printf("1/%.4f = %.4f\n", arg, result)
		case 3:
			// arccos 0.79
			result := math.Acos(arg)
			fmt.Printf("arccos(%.4f) = %.4f\n", arg, result)
		case 4:
			// e^2.01
			result := math.Exp(arg)
			fmt.Printf("e^(%.4f) = %.4f\n", arg, result)
		case 5:
			// 3.4^2.6
			result := math.Pow(3.4, arg)
			fmt.Printf("3.4^(%.4f) = %.4f\n", arg, result)
		}
	}
}
