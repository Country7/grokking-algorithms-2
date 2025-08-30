package main

// Алгоритм Дейкстры - поиск кратчайшего пути во взвешенных графах с положительной стоимостью ребер
// Временная сложность алгоритма O(n²), где n — количество вершин.

import (
	"fmt"
	"math"
	"strings"
)

// Опишем граф со всеми узлами и взвешенными ребрами из упражения 9.1 А
var graph = map[string]map[string]float64{
	"start": {
		"A": 5,
		"B": 2,
	},
	"A": {
		"C": 4,
		"D": 2,
	},
	"B": {
		"A": 8,
		"D": 7,
	},
	"C": {
		"D":   6,
		"fin": 3,
	},
	"D": {
		"fin": 1,
	},
	"fin": {},
}

// Печатаем граф красиво
func printBeautiGraph() {
	for from, edges := range graph {
		fmt.Printf("  %s ─> ", from) // родительский узел ->
		if len(edges) == 0 {
			fmt.Println("∅")
			continue
		}
		for to, weight := range edges {
			fmt.Printf("%s(%.0f) ", to, weight) // к какому узлу (стоимость)
		}
		fmt.Println()
	}
}

// Основная функция алгоритма Дейкстры
func dijkstrasAlgoritm(start, finish string) ([]string, float64) {

	// Таблица стоимостей ребер (известное расстояние от старта до каждой вершины)
	var costs = make(map[string]float64)

	// Родители (для восстановления пути)
	var parents = make(map[string]string)

	// Множество обработанных узлов
	var processed = make(map[string]bool)

	// Инициализация: все вершины = ∞
	for node := range graph {
		costs[node] = math.Inf(1) // coast map [start:+Inf a:+Inf b:+Inf fin:+Inf]
		parents[node] = ""        // parents map [a: b: fin: start:]
	}
	// Для стартовой вершины вес равен 0
	costs[start] = 0 // coast map [start:0 a:+Inf b:+Inf fin:+Inf]
	fmt.Println("\nИнициализация всех вершин в ∞, стартовой в 0: costs =", costs, "\n ")

	// Пока есть необработанные узлы перебираем их
	for {
		// Найти узел с наименьшей стоимостью, если есть хоть один не пустой - продолжаем цикл
		node := findLowestCostNode(costs, processed)
		fmt.Println("───> Узел с наименьшей стоимостью", node)
		if node == "" {
			break
		}

		// Обновить стоимости для соседей
		cost := costs[node] // Стоимость узла до node
		fmt.Printf("  ╰──> Обновляем стоимости для соседей %s(%.0f). ", node, cost)
		fmt.Println("Список соседей и вес их ребер ->", graph[node]) // из исходного графа
		fmt.Print("    ╰──> ")
		for n, weight := range graph[node] {
			fmt.Printf("Было -> %s(%.0f) ", n, costs[n]) // Какая стоимость узла была
			newCost := cost + weight
			if newCost < costs[n] {
				costs[n] = newCost
				parents[n] = node
			}
			fmt.Printf("Стало -> %s(%.0f);  ", n, costs[n]) // Какая стоимость узла стала
		}
		fmt.Println()
		// Пометить узел как обработанный
		processed[node] = true
	}

	// Путь найден, Восстановление пути
	path := []string{}
	curr := finish
	fmt.Printf("\nВосстановление пути: %s", curr)
	for curr != "" {
		path = append([]string{curr}, path...)
		curr = parents[curr]
		fmt.Printf(", %s", curr)
	}
	fmt.Println("\n  ╰──> Путь:", path, "\n ")

	return path, costs[finish]
}

// Вспомогательная функция: найти узел с минимальной стоимостью
func findLowestCostNode(costs map[string]float64, processed map[string]bool) string {
	lowestCost := math.Inf(1)
	lowestNode := ""
	for node, cost := range costs {
		if cost < lowestCost && !processed[node] {
			lowestCost = cost
			lowestNode = node
		}
	}
	return lowestNode
}

// Визуализация графа в ASCII
func printGraph(path []string) {
	fmt.Println("Граф ( * помечены ребра, которые входят в кратчайший путь):")
	for from, edges := range graph {
		for to, w := range edges {
			edge := fmt.Sprintf("  %s ─(%.0f)─> %s ", from, w, to)
			if isEdgeInPath(from, to, path) {
				edge += "  *" // помечаем рёбра кратчайшего пути
			}
			fmt.Println(edge)
		}
	}
}

// Проверка, есть ли ребро в найденном пути
func isEdgeInPath(from, to string, path []string) bool {
	for i := 0; i < len(path)-1; i++ {
		if path[i] == from && path[i+1] == to {
			return true
		}
	}
	return false
}

// ─────────────╯

// ─────────────╮
func main() {
	fmt.Println("\nАЛГОРИТМ ДЕЙКСТРЫ \n\nИсходный граф: graph =")
	printBeautiGraph()

	path, cost := dijkstrasAlgoritm("start", "fin")

	fmt.Printf("Кратчайший путь: %s\n", strings.Join(path, " -> "))
	fmt.Printf(" ╰──> Стоимость: %.0f\n\n", cost)
	printGraph(path)
}
