package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const todoFile = "todos.txt"

// Todo represents a single task
type Todo struct {
	ID   int
	Task string
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task to add")
			printUsage()
			os.Exit(1)
		}
		task := strings.Join(os.Args[2:], " ")
		addTodo(task)
	case "list":
		listTodos()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID to delete")
			printUsage()
			os.Exit(1)
		}
		id, err := parseID(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		deleteTodo(id)
	default:
		fmt.Println("Error: Unknown command")
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: go run main.go <command> [args]")
	fmt.Println("Commands:")
	fmt.Println("  add <task>    - Add a new task")
	fmt.Println("  list          - List all tasks")
	fmt.Println("  delete <id>   - Delete a task by ID")
	fmt.Println("Example: go run main.go add \"Buy groceries\"")
}

func loadTodos() ([]Todo, error) {
	file, err := os.OpenFile(todoFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var todos []Todo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		id, err := parseID(parts[0])
		if err != nil {
			continue
		}
		todos = append(todos, Todo{ID: id, Task: parts[1]})
	}
	return todos, scanner.Err()
}

func saveTodos(todos []Todo) error {
	file, err := os.Create(todoFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, todo := range todos {
		fmt.Fprintf(writer, "%d:%s\n", todo.ID, todo.Task)
	}
	return writer.Flush()
}

func addTodo(task string) {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	// Find the next available ID
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	newID := maxID + 1

	todos = append(todos, Todo{ID: newID, Task: task})
	if err := saveTodos(todos); err != nil {
		fmt.Println("Error saving todos:", err)
		os.Exit(1)
	}
	fmt.Printf("Added task: %d: %s\n", newID, task)
}

func listTodos() {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("Todo List:")
	for _, todo := range todos {
		fmt.Printf("%d: %s\n", todo.ID, todo.Task)
	}
}

func deleteTodo(id int) {
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	updatedTodos := []Todo{}
	found := false
	for _, todo := range todos {
		if todo.ID == id {
			found = true
		} else {
			updatedTodos = append(updatedTodos, todo)
		}
	}

	if !found {
		fmt.Printf("Error: Task with ID %d not found\n", id)
		os.Exit(1)
	}

	if err := saveTodos(updatedTodos); err != nil {
		fmt.Println("Error saving todos:", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted task with ID %d\n", id)
}

func parseID(idStr string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(idStr))
}
