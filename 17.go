package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res, count := binarySearch(arr, 1)
	fmt.Println(res, count)
}

func binarySearch(arr []int, search int) (result int, searchCount int) {
	mid := len(arr) / 2
	switch {
	case len(arr) == 0:
		fmt.Printf("%d - Not Found", search)
		return -1, searchCount
	case arr[mid] > search:
		result, searchCount = binarySearch(arr[:mid], search)
	case arr[mid] < search:
		result, searchCount = binarySearch(arr[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: //
		result = mid
	}
	searchCount++
	return
}
