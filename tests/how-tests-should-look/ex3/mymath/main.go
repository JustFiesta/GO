// Provide basic mathematical functions - this will be doc overwiew
package mymath

import "sort"

// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values. - this is an function description
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}