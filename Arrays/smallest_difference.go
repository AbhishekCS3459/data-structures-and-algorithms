/*

	This code implements the Smallest Difference problem which takes two arrays of integers as input and returns a pair of integers,
	one from each array, with the smallest absolute difference between them.

	The function first initializes two variables current and smallest to the maximum integer value. It then sorts both input arrays
	in ascending order using the sort.Ints function from the sort package.

	The function then iterates through both arrays using two pointers, idx1 and idx2, initialized to 0. Inside the loop, it compares
	the elements at the current indices of the two arrays, first and second, and calculates the absolute difference between
	them in the current variable.

	If current is smaller than the smallest variable, it updates smallest to current and assigns the current pair of integers
	to the result variable.

	The function returns the result variable, which contains the pair of integers with the smallest absolute difference.

	If there are identical integers in the two input arrays, the function will return them immediately, without any further comparisons.
*/
package main

import (
	"math"
	"sort"
)

// SmallestDifference takes two integer slices as input and returns a slice with two integers.
// The two integers in the returned slice have the smallest absolute difference among all pairs
// of integers from the two input slices.
func SmallestDifference(array1, array2 []int) []int {
	// Initialize variables for the smallest difference and the current difference being calculated
	current, smallest := math.MaxInt32, math.MaxInt32
	// Sort the input slices
	sort.Ints(array1)
	sort.Ints(array2)
	// Initialize variables for the indices for the two slices
	idx1, idx2 := 0, 0
	// Initialize an empty slice for the result
	result := []int{}
	// Loop through the two slices until we reach the end of one of the slices
	for idx1 < len(array1) && idx2 < len(array2) {
		// Get the values at the current indices for the two slices
		first, second := array1[idx1], array2[idx2]
		// Calculate the current difference between the two values
		if first < second {
			current = second - first
			idx1++
		} else if second < first {
			current = first - second
			idx2++
		} else {
			// If the two values are equal, we can return the pair
			return []int{first, second}
		}
		// Update the smallest difference and result slice if the current difference is smaller
		if smallest > current {
			smallest = current
			result = []int{first, second}
		}
	}
	// Return the pair with the smallest absolute difference
	return result
}
