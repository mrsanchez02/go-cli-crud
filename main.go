package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mrsanchez02/go-cli-crud/task"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type your task, please.")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)
		task.ListTasks(tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a valid ID for delete.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid Id")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)
		task.ListTasks(tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a valid task ID to complete.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid Id")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)
		task.ListTasks(tasks)
	default:
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Use: go-cli-crud [list|add|complete|delete]")
}
