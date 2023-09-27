package main

import (
	"fmt"
)

func main() {
	// Define the given approximate numbers and their error limits
	numbers := []struct {
		value float64
		error float64
	}{
		{42753.8, 800},
		{63.2561, 0.001},
		{63.2561, 0.002},
		{42753.8, 100},
		{2.53, 0.00001},
		{42153.8, 800},
		{2.53, 0.00004},
		{42743.8, 100},
	}

	// Iterate through the numbers and calculate absolute and relative errors
	for i, num := range numbers {
		absoluteError := num.error
		relativeError := (num.error / num.value) * 100

		// Round the relative error to a reasonable number of decimal places
		roundedRelativeError := round(relativeError, 4)

		fmt.Printf("For approximate number %d:\n", i+1)
		fmt.Printf("Value: %.4f\n", num.value)
		fmt.Printf("Absolute Error: %.4f\n", absoluteError)
		fmt.Printf("Relative Error: %.4f%%\n", roundedRelativeError)
		fmt.Println()
	}
}

// Round a float64 number to the specified number of decimal places
func round(num float64, decimalPlaces int) float64 {
	rounding := 1.0
	for i := 0; i < decimalPlaces; i++ {
		rounding *= 10.0
	}
	return float64(int((num*rounding)+0.5)) / rounding
}
