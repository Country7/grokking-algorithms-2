package main

import "fmt"

func checkBin(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == item {
			return mid
		}
		if item > list[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{1, 3, 5, 7, 9, 11, 15, 17}
	fmt.Println("Поиск в срезе:", list, "значения 3. Ответ - позиция", checkBin(list, 3)) // 1
	fmt.Println("Поиск в срезе:", list, "значения 8. Ответ - позиция", checkBin(list, 8)) // -1
}
