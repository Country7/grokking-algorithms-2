package main

import "fmt"

// Находим наименьшее значение в срезе
func findSmallest(slice []int) int {
	smallest := slice[0] // сохраняем наименьшее значение
	smallest_index := 0  // сохраняем индекс наименьшего значения
	for i := 0; i < len(slice); i++ {
		if slice[i] < smallest { // находим и сохраняем наименьшее значение и индекс
			smallest = slice[i]
			smallest_index = i
		}
	}
	return smallest_index
}

// Сортируем срез
func sliceSortSmollest(slice []int) []int {
	lenSlice := len(slice)
	newSlice := make([]int, lenSlice) // создаем новый срез по длинне предыдущего и инициируем нулями
	for i := 0; i < lenSlice; i++ {
		smallest_index := findSmallest(slice)                               // находим индекс наименьшего значения в срезе
		newSlice[i] = slice[smallest_index]                                 // добавляем элемент в новый срез
		slice = append(slice[:smallest_index], slice[smallest_index+1:]...) // удаляем наименьшее из исходного среза
	}
	return newSlice
}

func main() {
	slice := []int{12, 10, 15, 17, 35, 9, 7, 3, 21, 48, 27, 15, 11, 33}
	fmt.Println("Сортируем слайс по убыванию: ", slice)
	fmt.Println("Сортированный слайс: ", sliceSortSmollest(slice))

}
