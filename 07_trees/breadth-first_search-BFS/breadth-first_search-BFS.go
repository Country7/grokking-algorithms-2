package main

// Алгоритм Поиска в ширину (BFS) (без рекурсии)

import (
	"fmt"
	"os"
	"path/filepath"
	// "sort"
)

func printNamesBFS(root string) {
	// очередь папок для обработки
	queue := []string{root}

	for len(queue) > 0 {
		// достаем первую папку из очереди
		dir := queue[0]
		queue = queue[1:]

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Ошибка чтения директории:", err)
			continue
		}

		/*
			// сортируем по имени
			sort.Slice(entries, func(i, j int) bool {
				return entries[i].Name() < entries[j].Name()
			})
		*/

		for _, entry := range entries {
			fullpath := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				// папку не обходим сразу, а добавляем в очередь
				queue = append(queue, fullpath)
			} else {
				// выводим имя файла
				fmt.Println(entry.Name())
			}
		}
	}
}

func main() {
	printNamesBFS("/Users/country/go/src/Grokking-algorithms-2/Training")
}
