package main

// Алгоритм Динамическое программирование
// Задача о рюкзаке в контексте собрать рюкзак в турпоход

import "fmt"

// Структура для хранения информации о предметах, необходимых в турпоход
type Object struct {
	name   string
	weight int // масса (вес) предмета в фунтах
	score  int // оценка (ценность, важность)
}

func main() {
	// Список предметов в турпоход
	objects := []Object{
		{"Вода", 3, 10},
		{"Книга", 1, 3},
		{"Еда", 2, 9},
		{"Куртка", 2, 5},
		{"Камера", 1, 6},
	}

	// Максимальная емкость рюкзака 6 футов
	maxWeight := 6
	n := len(objects)
	fmt.Println("\nЗадача взять самые нужные предметы в турпоход в рюкзак вместимостью", maxWeight, "футов:\n ")

	// Двумерная таблица DP: dp[i][t] = максимальная ценность,
	// которую можно получить, рассматривая первые i объектов и имея w веса
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxWeight+1)
	}

	// Заполняем таблицу динамического программирования
	for i := 1; i <= n; i++ {
		for w := 0; w <= maxWeight; w++ {
			if objects[i-1].weight <= w {
				// Либо берем предмет, либо нет
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-objects[i-1].weight]+objects[i-1].score)
			} else {
				// Если объема не хватает, не берем
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Находим максимальную длину имени предмета
	maxNameLen := 0
	for _, a := range objects {
		if len(a.name) > maxNameLen {
			maxNameLen = len(a.name)
		}
	}

	// Выводим таблицу DP
	fmt.Println("Таблица динамического программирования:")
	for i := -1; i < n; i++ {
		if i == -1 {
			// печатаем нулевую расчетную строку алгоритма
			fmt.Printf("%-*s", maxNameLen/2+18, "")
		} else {
			// Печатаем название в фиксированном поле (* — «ширина поля» передана как 1 аргумент)
			fmt.Printf("%-*s (t=%2d, s=%2d) -> ", maxNameLen/2+1, objects[i].name, objects[i].weight, objects[i].score)
		}
		// Печатаем значения dp
		for w := 0; w <= maxWeight; w++ {
			fmt.Printf("%3d ", dp[i+1][w]) // %3d = ширина ячейки 3 символа
		}
		fmt.Println()
	}

	// Восстановим оптимальный набор предметов
	resWeight := maxWeight
	chosen := []string{}
	for i := n; i > 0; i-- {
		if dp[i][resWeight] != dp[i-1][resWeight] {
			chosen = append(chosen, objects[i-1].name)
			resWeight -= objects[i-1].weight
		}
	}

	// Выводим результат
	fmt.Println("\nОптимальный набор предметов в рюкзаке вместимостью", maxWeight, "футов:")
	for i := len(chosen) - 1; i >= 0; i-- {
		fmt.Println("-", chosen[i])
	}
	fmt.Println("Максимальная суммарная ценность:", dp[n][maxWeight], "\n ")
}

// Вспомогательная функция: максимум из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
