package main

import (
	"fmt"
	"math"
)

func main() {
	x := 4.53
	relativeError := 0.02 / 100.0 // Convert relative error to decimal

	// Estimate the number of digits correct in the strict sense for each function
	digitsLnX := estimateCorrectDigits(math.Log(x), relativeError)
	digitsExpX := estimateCorrectDigits(math.Exp(x), relativeError)
	digitsPowXX := estimateCorrectDigits(math.Pow(x, x), relativeError)

	fmt.Printf("For x = %.2f with a relative error of %.4f%%:\n", x, relativeError*100)
	fmt.Printf("Estimated Correct Digits for ln(x): %d\n", digitsLnX)
	fmt.Printf("Estimated Correct Digits for e^x: %d\n", digitsExpX)
	fmt.Printf("Estimated Correct Digits for x^x: %d\n", digitsPowXX)
}

// Estimate the number of correct digits in the strict sense for a given value and relative error
func estimateCorrectDigits(value, relativeError float64) int {
	// Calculate the absolute error from the relative error
	absoluteError := relativeError * value

	// Calculate the number of correct digits using the absolute error
	correctDigits := int(math.Log10(1.0 / absoluteError))

	return correctDigits
}
