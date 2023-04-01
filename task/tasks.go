package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id       int    `json:"Id`
	Name     string `json:Name`
	Complete bool   `json:Complete`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks.")
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✔"
		}
		fmt.Printf("[%s] %d %s\n", status, task.Id, task.Name)
	}
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.Id == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Complete = !tasks[i].Complete
			break
		}
	}
	return tasks
}

func AddTask(tasks []Task, name string) []Task {
	newId := GetNextID(tasks)
	newTask := Task{
		Id:       newId,
		Name:     name,
		Complete: false,
	}
	tasks = append(tasks, newTask)
	return tasks
}

func SaveTasks(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}

func GetNextID(tasks []Task) int {
	amountTask := len(tasks)
	if amountTask == 0 {
		return 1
	}
	return tasks[len(tasks)-1].Id + 1
}
