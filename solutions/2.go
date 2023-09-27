package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []float64{3.009983, 24.00551, 21.161728}

	for _, num := range numbers {
		roundedTwoDecimals := round(num, 2)
		roundedThreeDecimals := round(num, 3)
		roundedFourDecimals := round(num, 4)

		fmt.Printf("Original Number: %.6f\n", num)
		fmt.Printf("Rounded to 2 Decimals: %.2f\n", roundedTwoDecimals)
		fmt.Printf("Rounded to 3 Decimals: %.3f\n", roundedThreeDecimals)
		fmt.Printf("Rounded to 4 Decimals: %.4f\n", roundedFourDecimals)
		fmt.Println()
	}
}

// Round a float64 number to the specified number of decimal places
func round(num float64, decimalPlaces int) float64 {
	shift := math.Pow(10, float64(decimalPlaces))
	return math.Round(num*shift) / shift
}
