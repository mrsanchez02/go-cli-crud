package task

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestListTasks(t *testing.T) {
	tasks := []Task{
		{Id: 1, Name: "Task 1", Complete: true},
		{Id: 2, Name: "Task 2", Complete: false},
	}

	ListTasks(tasks)
	// Output: [✔] 1 Task 1
	//         [ ] 2 Task 2
}

func TestDeleteTask(t *testing.T) {
	tasks := []Task{
		{Id: 1, Name: "Task 1", Complete: true},
		{Id: 2, Name: "Task 2", Complete: false},
		{Id: 3, Name: "Task 3", Complete: false},
	}

	tasks = DeleteTask(tasks, 2)

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, but got %d", len(tasks))
	}

	if tasks[0].Id != 1 {
		t.Errorf("Expected first task ID to be 1, but got %d", tasks[0].Id)
	}

	if tasks[1].Id != 3 {
		t.Errorf("Expected second task ID to be 3, but got %d", tasks[1].Id)
	}
}

func TestCompleteTask(t *testing.T) {
	tasks := []Task{
		{Id: 1, Name: "Task 1", Complete: false},
		{Id: 2, Name: "Task 2", Complete: false},
	}

	tasks = CompleteTask(tasks, 1)

	if tasks[0].Complete != true {
		t.Errorf("Expected first task to be complete, but got %t", tasks[0].Complete)
	}

	tasks = CompleteTask(tasks, 2)

	if tasks[1].Complete != true {
		t.Errorf("Expected second task to be complete, but got %t", tasks[1].Complete)
	}

	tasks = CompleteTask(tasks, 1)

	if tasks[0].Complete != false {
		t.Errorf("Expected first task to be incomplete, but got %t", tasks[0].Complete)
	}
}

func TestAddTask(t *testing.T) {
	tasks := []Task{
		{Id: 1, Name: "Task 1", Complete: true},
		{Id: 2, Name: "Task 2", Complete: false},
	}

	tasks = AddTask(tasks, "Task 3")

	if len(tasks) != 3 {
		t.Errorf("Expected 3 tasks, but got %d", len(tasks))
	}

	if tasks[2].Name != "Task 3" {
		t.Errorf("Expected third task name to be 'Task 3', but got '%s'", tasks[2].Name)
	}

	if tasks[2].Complete != false {
		t.Errorf("Expected third task to be incomplete, but got %t", tasks[2].Complete)
	}
}

func TestSaveTasks(t *testing.T) {
	// Create a temporary file to write the tasks to
	tmpfile, err := ioutil.TempFile("", "tasks.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create some tasks to save
	tasks := []Task{
		{Id: 1, Name: "Task 1", Complete: false},
		{Id: 2, Name: "Task 2", Complete: true},
		{Id: 3, Name: "Task 3", Complete: false},
	}

	// Save the tasks to the file
	SaveTasks(tmpfile, tasks)

	// Read the saved tasks from the file
	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		t.Fatal(err)
	}
	var savedTasks []Task
	err = json.Unmarshal(bytes, &savedTasks)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the saved tasks with the original tasks
	if len(savedTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, got %d", len(tasks), len(savedTasks))
	}
	for i := range tasks {
		if tasks[i].Id != savedTasks[i].Id {
			t.Errorf("Expected task %d ID %d, got ID %d", i+1, tasks[i].Id, savedTasks[i].Id)
		}
		if tasks[i].Name != savedTasks[i].Name {
			t.Errorf("Expected task %d name %q, got %q", i+1, tasks[i].Name, savedTasks[i].Name)
		}
		if tasks[i].Complete != savedTasks[i].Complete {
			t.Errorf("Expected task %d complete %v, got %v", i+1, tasks[i].Complete, savedTasks[i].Complete)
		}
	}
}

func TestGetNextID(t *testing.T) {
	// Test with an empty slice
	tasks := []Task{}
	nextID := GetNextID(tasks)
	if nextID != 1 {
		t.Errorf("Expected next ID to be 1, got %d", nextID)
	}

	// Test with a non-empty slice
	tasks = []Task{
		{Id: 1, Name: "Task 1", Complete: false},
		{Id: 2, Name: "Task 2", Complete: true},
		{Id: 3, Name: "Task 3", Complete: false},
	}
	nextID = GetNextID(tasks)
	if nextID != 4 {
		t.Errorf("Expected next ID to be 4, got %d", nextID)
	}
}
