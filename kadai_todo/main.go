package main

import (
	"fmt"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *Todo) Complete() {
	t.Completed = true
	fmt.Printf("ID:%dのTodoを完了に更新します\n", t.ID)
}

func printTools(t Todo) {
	if t.Completed {
		fmt.Printf("[完了] (ID: %d) %s\n", t.ID, t.Title)
	} else {
		fmt.Printf("[未完了] (ID: %d) %s\n", t.ID, t.Title)
	}
}

func main() {
	todos := []Todo{
		{ID: 1, Title: "学習計画", Completed: false},
		{ID: 2, Title: "環境構築", Completed: false},
		{ID: 3, Title: "基礎構文", Completed: false},
	}

	fmt.Println("=== Todoリスト ===")
	for i := range todos {
		printTools(todos[i])
	}

	todos[0].Complete() // ID:1のTodoを完了に更新します
	todos[1].Complete() // ID:2のTodoを完了に更新します

	fmt.Println("=== 更新後のTodoリスト ===")
	for i := range todos {
		printTools(todos[i])
	}
}
