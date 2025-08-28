package main

// Пример поиска в ширину и реализации графа
// Ищем ближайшего в окружении друзей продавца манго
// Временная сложность алгоритма: Поиск в ширину выполняется за время O(количество людей + количество ребер), что обычно записывается в форме O(V+E) (V — количество вершин, E — количество ребер)

import "fmt"

// var graph = make(map[string][]string) // Объявление и инициализация графа
var graph = map[string][]string{ // Объявление и инициализация и наполнение графа
	"you":    {"Alice", "Bob", "Claire"},
	"Bob":    {"Anuj", "Peggy"},
	"Alice":  {"Peggy"},
	"Claire": {"Thom", "Jonny"},
	"Anuj":   {},
	"Peggy":  {},
	"Thom":   {},
	"Jonny":  {},
}

func search(name string) bool {
	var search_queue []string                           // Создаем очередь
	search_queue = append(search_queue, graph[name]...) // Добавляем всех соседей в очередь
	fmt.Println("Первая очередь соседей:", search_queue)
	var searched []string // создаем список для отслеживания уже проверенных людей
	var person string     // создаем переменную под персоналия

	for len(search_queue) != 0 { // Пока очередь непустая
		person = search_queue[0]        // Извлекаем из очереди первую персону
		search_queue = search_queue[1:] // Убираем персону из очередь поиска

		if person_not_in_searched(person, searched) { // Человек проверяется только в том случае, если он не проверялся ранее, иначе может быть бесконечный цикл

			if person_is_seller(person) { // Проверяем, является ли этот человек продавцом манго
				fmt.Println(person, " продавец манго!") // Да, это продавец манго
				return true
			} else {
				search_queue = append(search_queue, graph[person]...) // Нет, не является. Все друзья этого человека добавляются в очередь поиска
				fmt.Println("Новая очередь соседей:", search_queue)
				searched = append(searched, person) // персона добавляется в список уже проверенных людей
			}
		}
	}

	return false // Если выполнение дошло до этой строки, значит, в очереди нет продавца манго
}

func person_is_seller(name string) bool {
	return name[len(name)-1] == 'm' // !!! в Go строковый литерал для одиночного символа должен быть РУНОЙ ('m')
	// Если используется обратная кавычка `m`, то в Go это raw string literal, строка, а не символ !!! Ошибка!
	// Эта функция проверяет, заканчивается ли имя на букву «m», и если заканчивается, этот человек считается продавцом манго. Проверка довольно глупая, но для нашего примера сойдет.
}

func person_not_in_searched(person string, searched []string) bool {
	for _, name := range searched {
		if person == name { // если человек уже проверялся
			return false
		}
	}
	return true // человек еще не проверялся
}

func main() {
	/* данные уже добавили при объявлении графа, но можно и так:
	graph["you"] = []string{"Alice", "Bob", "Claire"}
	graph["Bob"] = []string{"Anuj", "Peggy"}
	graph["Alice"] = []string{"Peggy"}
	graph["Claire"] = []string{"Thom", "Jonny"}
	graph["Anuj"] = []string{}
	graph["Peggy"] = []string{}
	graph["Thom"] = []string{}
	graph["Jonny"] = []string{}
	*/
	search("you")
}
