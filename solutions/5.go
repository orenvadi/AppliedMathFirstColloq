package main

import (
	"fmt"
	"math"
)

func main() {
	values := []float64{2.364, 109.6, 14.307}
	errors := []float64{0.07, 0.04, 0.005}

	for i, value := range values {
		// Calculate the number of correct digits in the strict sense
		correctDigitsStrict := countCorrectDigitsStrict(value, errors[i])

		// Round the value to the correct digits with one spare digit
		roundedValue := roundToCorrectDigits(value, correctDigitsStrict+1)

		fmt.Printf("Original Value %c:\n", 'a'+i)
		fmt.Printf("Value: %.4f\n", value)
		fmt.Printf("Relative Error: %.4f%%\n", errors[i])
		fmt.Printf("Number of Correct Digits (Strict Sense): %d\n", correctDigitsStrict)
		fmt.Printf("Rounded Value: %.4f\n", roundedValue)
		fmt.Println()
	}
}

// Count the number of correct digits in the strict sense based on the relative error
func countCorrectDigitsStrict(value, relativeError float64) int {
	// Calculate the absolute error from the relative error
	absoluteError := (relativeError / 100.0) * value

	// Calculate the number of correct digits using the absolute error
	correctDigits := int(math.Log10(1.0 / absoluteError))

	return correctDigits
}

// Round a float64 number to the specified number of digits
func roundToCorrectDigits(num float64, digits int) float64 {
	shift := math.Pow(10, float64(digits))
	return math.Round(num*shift) / shift
}
