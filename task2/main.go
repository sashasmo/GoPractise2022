package main

import (
	"fmt"
)

func BinarySearch(in []int, searchFor int) {
	if len(in) == 0 {
		fmt.Println("Ошибка. Массив пуст.")
	}

	var first, last = 0, len(in) - 1

	for first <= last {
		var mid = ((last - first) / 2) + first

		if in[mid] == searchFor {
			fmt.Println("Искомый элемент", searchFor, "имеет индекс", mid)
			break
		} else if in[mid] > searchFor { // нужно искать в левой части слайса
			last = mid - 1
		} else if in[mid] < searchFor { // нужно искать в правой части слайса
			first = mid + 1
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	BinarySearch(arr, 5)
}