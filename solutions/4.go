package main

import (
	"fmt"
)

func main() {
	numbers := []float64{0.310, 3.495, 24.3790}

	for _, num := range numbers {
		// Round the number to two decimal places
		rounded := roundToTwoDecimalPlaces(num)

		// Determine the number of digits that are correct in the strict sense
		correctDigitsStrict := countCorrectDigitsStrict(num, rounded)

		fmt.Printf("Original Number: %.4f\n", num)
		fmt.Printf("Rounded to Hundredths: %.2f\n", rounded)
		fmt.Printf("Number of Correct Digits (Strict Sense): %d\n", correctDigitsStrict)
		fmt.Println()
	}
}

// Round a float64 number to two decimal places
func roundToTwoDecimalPlaces(num float64) float64 {
	return float64(int(num*100+0.5)) / 100.0
}

// Count the number of correct digits in the strict sense between two numbers
func countCorrectDigitsStrict(original, rounded float64) int {
	// Convert the numbers to strings to compare digit by digit
	originalStr := fmt.Sprintf("%.2f", original)
	roundedStr := fmt.Sprintf("%.2f", rounded)

	// Initialize a count for correct digits
	count := 0

	// Iterate through the digits and count correct digits
	for i := 0; i < len(originalStr) && i < len(roundedStr); i++ {
		if originalStr[i] == roundedStr[i] {
			count++
		} else {
			break // Stop counting as soon as a digit is incorrect
		}
	}

	return count
}
