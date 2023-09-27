package main

import (
	"fmt"
)

func main() {
	operations := []struct {
		a, b float64
		op   string
	}{
		{1.504, 1.502, "-"},
		{12.64, 0.3, "-"},
		{24.37, 9.18, "-"},
		{18.437, 24.9, "+"},
		{234.5, 194.3, "-"},
		{72.3, 0.34, ":"},
		{24.1, 0.037, "-"},
		{0.65, 1984, "-"},
		{8124.6, 2.9, ":"},
	}

	for i, operation := range operations {
		var result float64
		switch operation.op {
		case "-":
			result = operation.a - operation.b
		case "+":
			result = operation.a + operation.b
		case ":":
			result = operation.a / operation.b
		}

		// Calculate the absolute uncertainty
		absoluteUncertainty := operation.a - operation.b

		// Calculate the relative uncertainty
		relativeUncertainty := (absoluteUncertainty / result) * 100

		fmt.Printf("For operation %c:\n", 'а'+i)
		fmt.Printf("Result: %.4f\n", result)
		fmt.Printf("Absolute Uncertainty: %.4f\n", absoluteUncertainty)
		fmt.Printf("Relative Uncertainty: %.4f%%\n", relativeUncertainty)
		fmt.Println()
	}
}
