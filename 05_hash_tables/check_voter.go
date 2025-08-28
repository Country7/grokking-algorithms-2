package main

import "fmt"

var voted map[string]bool

func check_voter(name string) {
	if voted[name] { // если имя в хеш-таблице есть
		fmt.Println(name, "- уже голосовал, на выход!") // значит уже проголосовал
	} else {
		voted[name] = true // иначе допускаем к голосованию и заносим в хеш-таблицу
		fmt.Println(name, "- пройдите для голосования!")
	}
}

func main() {
	voted = make(map[string]bool)
	check_voter("tom")
	check_voter("mike")
	check_voter("mike")
}
