package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для хранения тасков

type Todo struct {
	ID        int    `json:"ID"`       // Первичный ключ
	Task      string `json:"Task"`     // Сама таска
	Completed bool   `json:"Comleted"` // Состояние таски (выполнено или нет)
}

var todos = []Todo{} // Слайс структур
var nextID = 1       // Если todo новый, то таски начинаются с 1, если существующий, то nextID = maxID + 1

func main() {
	defer recover()
	loadFromFile()

	scanner := bufio.NewScanner(os.Stdin) //Сканер с буфером для хранения
	for {
		fmt.Println("")
		fmt.Println("-----------------------------------------")
		fmt.Println("Напишите номер операции для её выполнения")
		fmt.Println("-----------------------------------------")
		fmt.Println("1. Показать список задач")
		fmt.Println("2. Добавляение задачи")
		fmt.Println("3. Изменение состояние задачи")
		fmt.Println("4. Удаление задачи")
		fmt.Println("5. Сохранить и выйти")
		fmt.Println("6. TODO")
		fmt.Println("-----------------------------------------")

		scanner.Scan()           // Сканируем вводимую операцию
		choise := scanner.Text() // Переносим это в переменную choise

		switch choise { // чузим операцию
		case "1":
			showTodos()
		case "2":
			addTodo(scanner)
		case "3":
			completeTodo(scanner)
		case "4":
			deleteTodo(scanner)
		case "5":
			saveToFile()
			panic("123")
		default:
			fmt.Println("Неверная операция, попробуйте ещё раз")
		}

	}

}

func addTodo(scanner *bufio.Scanner) {
	fmt.Println("Введите задачу")
	scanner.Scan()
	task := strings.TrimSpace(scanner.Text()) // удаляем пробелы в начале и конце (если они есть)
	if task == "" {
		fmt.Println("Задача не может быть пустой")
		return
	}
	todo := Todo{
		ID:        nextID,
		Task:      task,
		Completed: false,
	}
	todos = append(todos, todo) // Обновляем существующий слайс
	nextID++                    // Меняем айди для создание новой задачи
	fmt.Println("Задача добавлена!")
}

func saveToFile() {
	data, err := json.MarshalIndent(todos, "", "  ") // Сохраняем не в одной строчке (для красоты)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
	os.WriteFile("todos.json", data, 0644)
}

func showTodos() {
	if len(todos) == 0 {
		fmt.Println("Задач нет")
		return
	}
	for _, todo := range todos {
		status := "❌"

		if todo.Completed {
			status = "✅"

		}
		fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Task)
	}

}

func completeTodo(scanner *bufio.Scanner) {
	fmt.Print("Введите ID задачи для завершения: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text()) //Конвертируем text в int

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			fmt.Println("Задача завершена!")
			return
		}
	}
	fmt.Println("Задача не найдена")
}
func deleteTodo(scanner *bufio.Scanner) {
	fmt.Print("Введите ID задачи для удаления: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text()) //Конвертируем text в int

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Задача удалена!")
			return
		}
	}
	fmt.Println("Задача не найдена")
}

func loadFromFile() {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		fmt.Println("Ошибка, нету JSON")
		return
	}

	json.Unmarshal(data, &todos)

	for _, todo := range todos {
		if todo.ID > nextID {
			nextID = todo.ID + 1
		}
	}
}
