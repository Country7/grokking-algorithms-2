package main

import "fmt"

func RecursiveCheckBin(list []int, item int, low, high int) int {
	if low <= high {
		mid := (low + high) / 2
		if list[mid] == item {
			return mid
		} else if item < list[mid] {
			return RecursiveCheckBin(list, item, low, mid-1) // возвращаем ф-цию с новыми значениями
		} else {
			return RecursiveCheckBin(list, item, mid+1, high) // возвращаем ф-цию с новыми значениями
		}
	}
	return -1
}

func main() {
	list := []int{1, 3, 5, 7, 9, 11, 15, 17}
	fmt.Println("Поиск в срезе:", list, "значения 3. Ответ - позиция", RecursiveCheckBin(list, 3, 0, len(list)-1)) // 1
	fmt.Println("Поиск в срезе:", list, "значения 8. Ответ - позиция", RecursiveCheckBin(list, 8, 0, len(list)-1)) // -1
}
