package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 7, 5, 9, 0, 2, 1}
	fmt.Printf("---Before--- %d\n", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("---After--- %d\n", arr)
}

// quickSort обертка рекурсивного вызова
func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot)
		quickSort(arr, pivot+1, high)
	}
}

// partition раскидывает элемента которые меньше опорного налево, которые больше направо
func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j > low {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], pivot
	return j
}
