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

func sort(data []int) []int {
	var n int = len(data)

	for i := 1; i < n; i++ {
		var key int = data[i]
		var j int = i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}

	return data
}

func reverseSort(data []int) []int {
	var n int = len(data)

	for i := 1; i < n; i++ {
		var key int = data[i]
		var j int = i - 1
		for j >= 0 && data[j] < key {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}

	return data
}

func visualize(data []int) {
	for _, val := range data {
		fmt.Println(val, createBar(val))
	}
}

func createBar(n int) string {
	var bar string = ""

	for i := 1; i <= n; i++ {
		bar += "-"
	}

	return bar
}
