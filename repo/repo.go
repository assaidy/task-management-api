package repo

import (
	"errors"
	"github.com/assaidy/task-manager-api/models"
	"sync"
)

// tasks stores the list of tasks
var tasks []models.Task

// currentId keeps track of the last used task ID
var currentId int

// ErrorNotFound is returned when a task with a specific ID is not found
var ErrorNotFound = errors.New("Task Not Found")

// mu is used to synchronize access to the tasks slice
var mu sync.Mutex

func init() {
    tasks = make([]models.Task, 0)
}

// GetAllTasks returns all the tasks
func GetAllTasks() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}

// CreateTask creates a new task with the given content
func CreateTask(content string) models.Task {
	mu.Lock()
	defer mu.Unlock()
	currentId++
	t := models.Task{Id: currentId, Content: content}
	tasks = append(tasks, t)
	return t
}

// GetTask returns a task by its ID
func GetTask(id int) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range tasks {
		if t.Id == id {
			return t, nil
		}
	}
	return models.Task{}, ErrorNotFound
}

// ToggleTaskComplete toggles the completion status of a task by its ID
func ToggleTaskComplete(id int) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Completed = !tasks[i].Completed
			return tasks[i], nil
		}
	}
	return models.Task{}, ErrorNotFound
}

// DeleteTask deletes a task by its ID
func DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return ErrorNotFound
}

// UpdateTask updates the content of a task by its ID
func UpdateTask(id int, newContent string) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.Id == id {
			tasks[i].Content = newContent
			return tasks[i], nil
		}
	}
	return models.Task{}, ErrorNotFound
}

// DeleteAllTasks deletes all tasks
func DeleteAllTasks() {
	mu.Lock()
	defer mu.Unlock()
	tasks = make([]models.Task, 0)
}
