package main

// Задача о покрытии множества. Подобрать радиостанции с макс покрытием
// Жадные алгоритмы, Сложность O(n^2)

import (
	"fmt"
)

// Множество радиостанций с покрытием по штатам
var stations = map[string][]string{
	"kone":   {"id", "nv", "ut"},
	"ktwo":   {"wa", "id", "mt"},
	"kthree": {"or", "nv", "ca"},
	"kfour":  {"nv", "ut"},
	"kfive":  {"ca", "az"},
}

// ─────────────╮
// Основная функция поиска станций
func findStations() []string {

	fmt.Println("\nМножество радиостанций с покрытием по штатам: \n", stations)

	// Множество всех имеющихся станций
	stationsKey := allStations()
	fmt.Println("\n  ╰──> Список всех возможных подмножеств станций:", stationsKey)

	// Множество штатов для покрытия радиостанциями
	states_needed := allStates()
	fmt.Println("\n  ╰──> Список всех штатов, которые необходимо покрыть:", states_needed, "\n ")

	// Итоговый набор станций
	var final_stations []string

	// Пока есть оставшиеся непокрытые радио штаты
	for len(states_needed) > 0 {
		var best_station string     // Лучшая станция по покрытию
		var states_covered []string // Все штаты, обслуживаемые этой станцией

		// Проходим циклом по множеству всех имеющихся станций
		for _, station := range stationsKey {
			states := stations[station] // берем штаты, обслуживаемые этой станцией
			// вычисляем пересечение необходимых штатов, и штатов ослуживаемых этой станцией
			covered := equaldata(states_needed, states)

			// Если пересеченное покрытие этой станции больше, чем покрытие лучшей станции
			if len(covered) > len(states_covered) {
				best_station = station   // тогда она становится лучшей
				states_covered = covered // и ее покрытие присваиваем states_covered
				fmt.Println("    ╰──> Лучшая станция в цикле:", best_station, "  и ее покрытие:", states_covered)
			}
		}
		// Удаляем из необбходимых штатов штаты, покрываемые лучшей станцией
		states_needed = removedata(states_needed, states_covered)
		fmt.Println("  ──> Теперь список штатов, необходимых для покрыти:", states_needed)

		// Добавляем в список отобранных станций лучшую
		final_stations = append(final_stations, best_station)
		fmt.Println("  ──> Теперь список отобранных станций:", states_needed, "\n ")
	}

	// РЕЗУЛЬТАТ выполнения задачи
	return final_stations
}

// Дополнительная функция выдает Множество всех имеющихся станций
func allStations() []string {
	var stationsKey []string
	for key, _ := range stations {
		stationsKey = append(stationsKey, key)
	}
	return stationsKey
}

// Дополнительная функция выдает Множество всех штатов
func allStates() []string {
	var tempMap = make(map[string]struct{})
	var listStates []string
	for _, stationStates := range stations {
		for _, state := range stationStates {
			if _, ok := tempMap[state]; !ok {
				tempMap[state] = struct{}{}
				listStates = append(listStates, state)
			}
		}
	}
	return listStates
}

// Дополнительная функция для поиска пересечений необходимых штатов, и штатов ослуживаемых станцией
func equaldata(states_needed []string, states []string) []string {
	var covered []string
	for _, oneState_needed := range states_needed {
		for _, oneState := range states {
			if oneState_needed == oneState {
				covered = append(covered, oneState_needed)
			}
		}
	}
	return covered
}

// Дополнительная функция - Удаляем из необбходимых штатов штаты, покрываемые лучшей станцией
func removedata(states_needed []string, states_covered []string) []string {
	for _, oneState_covered := range states_covered { // Проходим циклом по всем штатам, покрываемым лучшей станцией
		for i, oneState_needed := range states_needed { // Проходим циклом по всем штатам, нужным для покрытия
			if oneState_covered == oneState_needed { // Если штат станции == штату со всех нужных штатов
				states_needed = append(states_needed[:i], states_needed[i+1:]...) // удаляем его из среза всех нужных шатов
			}
		}
	}
	return states_needed
}

// ─────────────╯

// ─────────────╮
func main() {
	fmt.Println("\nРЕШЕНИЕ: минимальное количество радиостанций с максимальным покрытием:", findStations(), "\n ")
}
