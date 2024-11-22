package services

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var tasks []*Task

func init() {
	loadTasks()
}

func loadTasks() error {
	filePath := "tasks/tasks.json"

	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	if err = json.NewDecoder(f).Decode(&tasks); err != nil {
		return fmt.Errorf("unable to decode JSON: %v", err)
	}

	return nil
}

func AddTask(description string) (int, error) {
	newTask := Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if len(tasks) > 0 {
		newTask.Id = tasks[len(tasks)-1].Id + 1
	} else {
		newTask.Id = 1
	}

	tasks = append(tasks, &newTask)

	return newTask.Id, saveTasks()
}

func saveTasks() error {
	filePath := "tasks/tasks.json"
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer f.Close()

	// if err := json.NewEncoder(f).Encode(tasks); err != nil {
	// 	return fmt.Errorf("unable to marshal tasks: %v", err)
	// }

	datas, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return fmt.Errorf("unable to marshal tasks: %v", err)
	}
	_, err = f.Write(datas)
	if err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}

func UpdateTask(id int, description string) error {
	for _, task := range tasks {
		if task.Id == id {
			task.Description = description
			task.UpdatedAt = time.Now()
			return saveTasks()
		}
	}
	return fmt.Errorf("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks()
		}
	}

	return fmt.Errorf("task not found")
}

func UpdateProgress(id int, status string) error {
	for _, task := range tasks {
		if task.Id == id {
			task.Status = status
			task.UpdatedAt = time.Now()
			return saveTasks()
		}
	}

	return fmt.Errorf("task not found")
}

func ListTasks(status string) ([]*Task, error) {
	// var finalTasks []string
	var finalTasks []*Task

	if status == "all" {
		// for _, task := range tasks {
		// 	finalTasks = append(finalTasks, task.Description)
		// }
		return tasks, nil
	} else {
		for _, task := range tasks {

			if task.Status == status {
				finalTasks = append(finalTasks, task)

			}
		}
	}

	return finalTasks, nil
}
