package service

import (
	"encoding/json"
	task "example/cli-taskTracker/entities"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getStoragePath() string {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	return filepath.Join(dir, "../storage/storage.json")
}

func ReadStorage() ([]byte, error) {
	path := getStoragePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UnmarshalStorage(data []byte) ([]task.Task, error) {
	var tasks []task.Task
	err := json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetId(tasks []task.Task) (int, error) {
	length := len(tasks)
	if length == 0 {
		return 1, nil
	}

	return tasks[length-1].Id + 1, nil
}

func AddTask(args []string, tasks []task.Task, id int) error {
	path := getStoragePath()
	newTask := task.NewTask(id, args[2])

	tasks = append(tasks, newTask)

	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error during marshalling tasks structure: %w", err)
	}

	return os.WriteFile(path, b, 0644)
}

func UpdateTask(findId int, updatedDesc string, tasks *[]task.Task) error {
	path := getStoragePath()
	if findId < 0 || findId > len(*tasks) {
		return fmt.Errorf("no such index to update: %d", findId)
	}

	(*tasks)[findId-1].Description = updatedDesc
	(*tasks)[findId-1].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error during marshalling in updating task: %w", err)
	}

	return os.WriteFile(path, b, 0644)
}

func DeleteTask(findId int, tasks *[]task.Task) error {
	path := getStoragePath()
	if findId < 0 || findId > len(*tasks) {
		return fmt.Errorf("no such id to delete: %d", findId)
	}

	sl := []task.Task{}
	index := 1
	for _, v := range *tasks {
		if v.Id != findId {
			v.Id = index
			index++

			sl = append(sl, v)
		}
	}

	tasks = &sl
	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error during marshalling tasks: %w", err)
	}

	return os.WriteFile(path, b, 0644)
} // учитывай что json не должен быть пустым, как минимум []

func MarkInProgress(findId int, tasks *[]task.Task) error {
	path := getStoragePath()
	if findId < 0 || findId > len(*tasks) {
		return fmt.Errorf("no such id to mark: %d", findId)
	}

	(*tasks)[findId-1].Status = "in-progress"

	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error during marshalling tasks: %w", err)
	}

	return os.WriteFile(path, b, 0644)
}

func MarkDone(findId int, tasks *[]task.Task) error {
	path := getStoragePath()
	if findId < 0 || findId > len(*tasks) {
		return fmt.Errorf("no such id to mark: %d", findId)
	}

	(*tasks)[findId-1].Status = "done"

	b, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("error during marshalling tasks: %w", err)
	}

	return os.WriteFile(path, b, 0644)
}

func PrintTasks(tasks []task.Task) {
	fmt.Printf(
		"%-4s | %-30s | %-12s | %-20s | %-20s\n",
		"ID", "DESCRIPTION", "STATUS", "CREATED AT", "UPDATED AT",
	)
	fmt.Println(strings.Repeat("-", 90))

	for _, t := range tasks {
		fmt.Printf(
			"%-4d | %-30s | %-12s | %-20s | %-20s\n",
			t.Id,
			t.Description,
			t.Status,
			t.CreatedAt,
			t.UpdatedAt,
		)
	}
}

func ShowList(tasks *[]task.Task) {
	PrintTasks(*tasks)
}

func ShowDone(tasks *[]task.Task) {
	sl := []task.Task{}
	for _, v := range *tasks {
		if v.Status == "done" {
			sl = append(sl, v)
		}
	}

	PrintTasks(sl)
}

func ShowToDo(tasks *[]task.Task) {
	sl := []task.Task{}
	for _, v := range *tasks {
		if v.Status == "todo" {
			sl = append(sl, v)
		}
	}

	PrintTasks(sl)
}

func ShowInProgress(tasks *[]task.Task) {
	sl := []task.Task{}
	for _, v := range *tasks {
		if v.Status == "in-progress" {
			sl = append(sl, v)
		}
	}

	PrintTasks(sl)
}

/*
task-cli:
add
update N
delete N

mark-in-progress N
mark-done N

list
list done
list todo
list in-progress
*/
