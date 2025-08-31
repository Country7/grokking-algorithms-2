package main

// Алгоритм Динамическое программирование
// Задача о рюкзаке в контексте туризма

import "fmt"

// Структура для хранения информации о достопримечательности
type Attraction struct {
	name  string
	time  int // время в условных единицах (0.5 дня = 1, 1 день = 2, 2 дня = 4)
	score int // оценка (ценность)
}

func main() {
	// Список достопримечательностей
	attractions := []Attraction{
		{"Вестминстерское аббатство", 1, 7},
		{"Театр 'Глобус'", 1, 6},
		{"Национальная галерея", 2, 9},
		{"Британский музей", 4, 9},
		{"Собор св. Павла", 1, 8},
		{"Центральная площадь", 1, 8},
		{"Кинотеатр", 1, 5},
	}

	// Общее доступное время: 3 дня = 6 половин дня
	maxTime := 6
	n := len(attractions)
	fmt.Println("\nЗадача посетить оптимальное число ценных достопримечательностей за", maxTime, "/ 2 дней:\n ")

	// Двумерная таблица DP: dp[i][t] = максимальная ценность,
	// которую можно получить, рассматривая первые i объектов и имея t времени
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxTime+1)
	}

	// Заполняем таблицу динамического программирования
	for i := 1; i <= n; i++ {
		for t := 0; t <= maxTime; t++ {
			if attractions[i-1].time <= t {
				// Либо берем достопримечательность, либо нет
				dp[i][t] = max(dp[i-1][t], dp[i-1][t-attractions[i-1].time]+attractions[i-1].score)
			} else {
				// Если времени не хватает, не берем
				dp[i][t] = dp[i-1][t]
			}
		}
	}

	// Находим максимальную длину имени достопримечательности
	maxNameLen := 0
	for _, a := range attractions {
		if len(a.name) > maxNameLen {
			maxNameLen = len(a.name)
		}
	}

	// Выводим таблицу DP
	fmt.Println("Таблица динамического программирования:")
	for i := -1; i < n; i++ {
		if i == -1 {
			// печатаем нулевую расчетную строку алгоритма
			fmt.Printf("%-*s", maxNameLen/2+16, "")
		} else {
			// Печатаем название в фиксированном поле (* — «ширина поля» передана как 1 аргумент)
			fmt.Printf("%-*s (t=%d, s=%d) -> ", maxNameLen/2+1, attractions[i].name, attractions[i].time, attractions[i].score)
		}
		// Печатаем значения dp
		for t := 0; t <= maxTime; t++ {
			fmt.Printf("%3d ", dp[i+1][t]) // %3d = ширина ячейки 3 символа
		}
		fmt.Println()
	}

	// Восстановим оптимальный маршрут
	resTime := maxTime
	chosen := []string{}
	for i := n; i > 0; i-- {
		if dp[i][resTime] != dp[i-1][resTime] {
			chosen = append(chosen, attractions[i-1].name)
			resTime -= attractions[i-1].time
		}
	}

	// Выводим результат
	fmt.Println("\nОптимальный маршрут:")
	for i := len(chosen) - 1; i >= 0; i-- {
		fmt.Println("-", chosen[i])
	}
	fmt.Println("Максимальная суммарная оценка:", dp[n][maxTime], "\n ")
}

// Вспомогательная функция: максимум из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
