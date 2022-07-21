package main

import "fmt"

func main() {
	var sample []int = []int{1, 4, 5, 6, 8, 2}

	fmt.Println("before sort visualization")
	visualize(sample)
	fmt.Println("before sort: ", sample)
	fmt.Println("ascending sort: ", sort(sample))
	fmt.Println("descending sort: ", reverseSort(sample))
}

// sort returns sorted slice in ascending
func sort(data []int) []int {
	// define the number of items inside slice
	var n int = len(data)

	// iterate through every item inside slice
	for i := 1; i < n; i++ {
		// define the temporary variable
		var key int = data[i]
		// define the additional index
		// for swapping
		var j int = i - 1

		// while j is greater than zero and
		// the item at j index is greater than temporary variable
		for j >= 0 && data[j] > key {
			// swap the data
			data[j+1] = data[j]
			j = j - 1
		}
		// assign the temporary variable
		data[j+1] = key
	}

	// return the sorted slice
	return data
}

// reverseSort returns sorted slice in descending
func reverseSort(data []int) []int {
	// define the number of items inside slice
	var n int = len(data)

	// iterate through every item inside slice
	for i := 1; i < n; i++ {
		// define the temporary variable
		var key int = data[i]
		// define the additional index
		// for swapping
		var j int = i - 1

		// while j is greater than zero and
		// the item at j index is less than temporary variable
		for j >= 0 && data[j] < key {
			// swap the data
			data[j+1] = data[j]
			j = j - 1
		}
		// assign the temporary variable
		data[j+1] = key
	}

	// return the sorted slice
	return data
}

// visualize create a visualization for slice data
func visualize(data []int) {
	// iterate through every item inside slice
	for _, val := range data {
		// create a bar chart
		fmt.Println(val, createBar(val))
	}
}

// createBar returns bar chart
func createBar(n int) string {
	// define the bar chart
	var bar string = ""

	// draw a bar chart based on the "n"
	for i := 1; i <= n; i++ {
		bar += "-"
	}

	// return the bar chart
	return bar
}
