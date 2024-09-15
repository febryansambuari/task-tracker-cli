package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func isValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	return validStatuses[status]
}

func readTasks(fileName string) ([]Task, error) {
	var tasks []Task

	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func writeTasks(fileName string, tasks []Task) error {
	updatedFile, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, updatedFile, 0644)
}

func addTask(descriptionTask string) {
	fileName := "tasks.json"

	tasks, err := readTasks(fileName)

	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading file:", err)
		return
	}

	newId := 1
	if len(tasks) > 0 {
		newId = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          newId,
		Description: descriptionTask,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	err = writeTasks(fileName, tasks)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Printf("Task added successfully with the ID: %d\n", newId)
}

func updateTask(taskId int, updatedDescriptionTask string) {
	fileName := "tasks.json"

	tasks, err := readTasks(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Find the task by ID
	var taskFound bool
	for i, task := range tasks {
		if task.ID == taskId {
			// Update task fields
			tasks[i].Description = updatedDescriptionTask
			tasks[i].UpdatedAt = time.Now()
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID: %d was not found\n", taskId)
		return
	}

	err = writeTasks(fileName, tasks)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	fmt.Printf("Task with ID %d has been updated.\n", taskId)
}

func deleteTask(taskId int) {
	fileName := "tasks.json"

	tasks, err := readTasks(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// Find the task by ID and delete it
	var taskFound bool
	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID: %d was not found\n", taskId)
		return
	}

	err = writeTasks(fileName, tasks)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	fmt.Printf("Task with ID: %d has been deleted.\n", taskId)
}

func markTask(taskId int, statusTask string) {
	fileName := "tasks.json"

	tasks, err := readTasks(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	if !isValidStatus(statusTask) {
		fmt.Printf("Invalid status: %s. Valid statuses are: todo, in-progress, done.\n", statusTask)
		return
	}

	var taskFound bool
	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].Status = statusTask
			tasks[i].UpdatedAt = time.Now()
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID: %d was not found\n", taskId)
		return
	}

	err = writeTasks(fileName, tasks)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	fmt.Printf("Task with ID: %d has been updated to: %s.\n", taskId, statusTask)
}

func listTasks(statusFilter string) {
	fileName := "tasks.json"

	tasks, err := readTasks(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No tasks found. The list is empty.")
			return
		}

		fmt.Println("Error reading file: ", err)
		return
	}

	// Check if there are any tasks
	if len(tasks) == 0 {
		fmt.Println("No tasks found. The list is empty.")
		return
	}

	// Display the list of tasks
	fmt.Println("Task List:")
	for _, task := range tasks {
		if statusFilter == "" || task.Status == statusFilter {
			createdAt := task.CreatedAt.Format("2006-01-02 15:04:05")
			updatedAt := task.UpdatedAt.Format("2006-01-02 15:04:05")

			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Printf("CreatedAt: %s\n", createdAt)
			fmt.Printf("UpdatedAt: %s\n", updatedAt)
			fmt.Println("---------------------------")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [command: add | update | delete | list | mark]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add [your description task]")
			return
		}
		descriptionTask := os.Args[2]
		addTask(descriptionTask)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go update [task_id] [your edited description task]")
			return
		}
		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		updatedDescriptionTask := os.Args[3]
		updateTask(taskId, updatedDescriptionTask)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete [task_id]")
			return
		}
		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		deleteTask(taskId)
	case "mark":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go mark [task_id] [status: todo | in-progress | done]")
			return
		}
		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		statusTask := os.Args[3]
		markTask(taskId, statusTask)
	case "list":
		var statusFilter string
		if len(os.Args) > 2 {
			statusFilter = os.Args[2]
		}
		listTasks(statusFilter)
	default:
		fmt.Println("Usage: go run main.go [command]")
	}
}
