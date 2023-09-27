package main

import (
	"fmt"
	"math"
)

func main() {
	// Define the values with all digits correct in the strict sense
	a := 0.62
	b := 16.9
	c := 41.3

	d := 12.47
	e := 12.5
	f := 14.8
	g := 0.97
	h := 2.63

	i := 6.91
	j := 3.35

	k := 26.88
	l := 3.94
	m := 8.04

	n := 6.19
	o := 1.34

	// Calculate the results
	resultA := (a + math.Sqrt(b)) / math.Log10(c)
	resultB := (d + math.Sqrt(e*e+f*f)) / (math.Sin(g)*math.Sin(g) + math.Cos(h)*math.Cos(h))
	resultC := math.Log(i+j*j) / math.Sqrt(626.3)
	resultD := math.Cbrt(k)/(math.Exp(l)-m*m) + math.Pow(n, o)

	// Print the results
	fmt.Printf("Result for a): %.4f\n", resultA)
	fmt.Printf("Result for b): %.4f\n", resultB)
	fmt.Printf("Result for c): %.4f\n", resultC)
	fmt.Printf("Result for d): %.4f\n", resultD)
}
