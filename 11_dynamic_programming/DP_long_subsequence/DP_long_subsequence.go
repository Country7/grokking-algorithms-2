package main

import (
	"fmt"
	"slices"
)

// Алгоритм Динамическое программирование
// Задача о поиске самой длинной общей подстроки и самой длинной общей последовательности

var words = []struct {
	a string
	b string
}{
	{a: "blue", b: "clues"},
	{a: "hish-vista", b: "hish-fish"},
}

// Доп. функция: создаём матрицу (двумерный массив) для хранения длин общих суффиксов
func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}
	return cell
}

// Функция находит самую длинную общую подстроку (Longest Common Substring)
// между двумя строками a и b и возвращает её.
func substring(a, b string) string {
	// lcs хранит длину найденной на данный момент максимальной общей подстроки
	lcs := 0
	// lastSubIndex хранит индекс в строке a, на котором заканчивается найденная общая подстрока
	lastSubIndex := 0
	// создаём матрицу (двумерный массив) для хранения длин общих суффиксов
	// cell[i][j] будет хранить длину наибольшей общей подстроки,
	// которая оканчивается в a[i-1] и b[j-1]
	cell := createMatrix(len(a)+1, len(b)+1)
	// Печатаем заголовок таблицы
	fmt.Print(" ")
	for j := 0; j < len(b); j++ {
		fmt.Printf("%3s", string(b[j]))
	}
	fmt.Println()
	// проходим по всем символам строк a и b
	for i := 1; i <= len(a); i++ {
		fmt.Printf("%-s", string(a[i-1]))
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] { // если символы совпадают
				// увеличиваем значение из "диагональной ячейки", то есть продолжаем подстроку
				cell[i][j] = cell[i-1][j-1] + 1
				if cell[i][j] > lcs { // если текущая длина больше найденного ранее максимума
					lcs = cell[i][j] // обновляем максимальную длину
					lastSubIndex = i // и сохраняем индекс конца подстроки в строке a
				}
			} else {
				cell[i][j] = 0 // если символы не совпадают — общей подстроки нет
			}
			fmt.Printf("%3d", cell[i][j])
		}
		fmt.Println()
	}
	// извлекаем подстроку из строки a
	// начиная с позиции lastSubIndex - lcs и до lastSubIndex
	return a[lastSubIndex-lcs : lastSubIndex]
}

// Функция возвращает длину самой длинной общей подпоследовательности
// (Longest Common Subsequence, LCS) между словами.
func subsequence(a, b string) (int, string) {
	// Жранит идексы совпадений символов
	var subIdxs []int
	// Хранит последовательность
	subSeq := ""
	// Создаем матрицу для динамического программирования
	cell := createMatrix(len(a)+1, len(b)+1)
	// Печатаем заголовок таблицы
	fmt.Print(" ")
	for j := 0; j < len(b); j++ {
		fmt.Printf("%3s", string(b[j]))
	}
	fmt.Println()
	// Два вложенных цикла перебирают все символы строк a и b.
	// Индексация начинается с 1, потому что 0-й ряд и 0-й столбец матрицы
	// служат "нулевым" базовым случаем (для пустой строки)
	for i := 1; i <= len(a); i++ {
		fmt.Printf("%-s", string(a[i-1]))
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				// Если символы a[i-1] и b[j-1] совпадают, значит текущая общая подпоследовательность
				// увеличивается на 1 по сравнению с cell[i-1][j-1]
				cell[i][j] = cell[i-1][j-1] + 1
				subIdxs = append(subIdxs, i-1)
			} else {
				// Если символы не совпадают, берется максимум из двух вариантов:
				// cell[i-1][j] — LCS без текущего символа a[i-1].
				// cell[i][j-1] — LCS без текущего символа b[j-1].
				cell[i][j] = cell[i-1][j]
				if cell[i][j] < cell[i][j-1] {
					cell[i][j] = cell[i][j-1]
				}
			}
			fmt.Printf("%3d", cell[i][j])
		}
		fmt.Println()
	}
	// Формируем строку последовательности
	if len(subIdxs) != 0 {
		for i := subIdxs[0]; i <= subIdxs[len(subIdxs)-1]; i++ {
			if slices.Contains(subIdxs, i) {
				subSeq += string(a[i])
			} else {
				subSeq += "."
			}
		}
	}

	// В последней ячейке матрицы хранится длина самой длинной
	// общей подпоследовательности для всей строки a и всей строки b.
	return cell[len(a)][len(b)], subSeq
}

func main() {
	fmt.Println("\nЗадача о поиске самой длинной общей подстроки и самой длинной общей последовательности\n ")
	fmt.Printf("Ищем подстроки в %s и %s \n", words[0].a, words[0].b)
	fmt.Printf("Самая длинная общая подстрока: %s \n \n", substring(words[0].a, words[0].b))

	fmt.Printf("Ищем последовательности в %s и %s \n", words[1].a, words[1].b)
	length, sub := subsequence(words[1].a, words[1].b)
	fmt.Println("Самая длинная общая последовательность:", sub, " длина:", length, "\n ")
}
