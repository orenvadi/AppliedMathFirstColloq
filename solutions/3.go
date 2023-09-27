package main

import (
	"fmt"
)

func main() {
	numbers := []float64{36.7, 2.489, 31.010, 0.031}

	for _, num := range numbers {
		// Calculate the limits of absolute errors
		absoluteErrorBroad := 0.5   // Broad sense absolute error is half of the smallest unit
		absoluteErrorStrict := 0.05 // Strict sense absolute error is one-tenth of the smallest unit

		// Calculate the limits of relative errors
		relativeErrorBroad := (absoluteErrorBroad / num) * 100
		relativeErrorStrict := (absoluteErrorStrict / num) * 100

		fmt.Printf("For number %.3f:\n", num)
		fmt.Printf("Broad Sense - Absolute Error Limit: %.3f\n", absoluteErrorBroad)
		fmt.Printf("Broad Sense - Relative Error Limit: %.3f%%\n", relativeErrorBroad)
		fmt.Printf("Strict Sense - Absolute Error Limit: %.3f\n", absoluteErrorStrict)
		fmt.Printf("Strict Sense - Relative Error Limit: %.3f%%\n", relativeErrorStrict)
		fmt.Println()
	}
}
