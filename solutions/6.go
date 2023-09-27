package main

import (
	"fmt"
	"math"
)

func main() {
	values := []float64{193, math.Sin(0.9), math.Log(24.6)}
	maxRelativeError := 0.1 // Maximum allowed relative error (0.1%)

	for i, value := range values {
		decimalPlaces := calculateDecimalPlaces(value, maxRelativeError)

		fmt.Printf("For value %c:\n", 'a'+i)
		fmt.Printf("Value: %.4f\n", value)
		fmt.Printf("Maximum Allowed Relative Error: %.4f%%\n", maxRelativeError)
		fmt.Printf("Number of Decimal Places (Strict Sense): %d\n", decimalPlaces)
		fmt.Println()
	}
}

// Calculate the number of decimal places required to achieve the specified relative error
func calculateDecimalPlaces(value, maxRelativeError float64) int {
	// Calculate the smallest absolute error that corresponds to the maximum relative error
	smallestAbsoluteError := (maxRelativeError / 100.0) * value

	// Calculate the number of decimal places needed to represent the smallest absolute error
	decimalPlaces := int(math.Ceil(-math.Log10(smallestAbsoluteError)))

	return decimalPlaces
}
