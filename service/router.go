package service

import (
	"example/cli-taskTracker/cli"
	"fmt"
	"log"
	"strconv"
)

func Router() {
	args := cli.GetInput()
	data, err := ReadStorage() // getting []byte from readin json
	if err != nil {
		log.Fatal(err)
	}

	tasks, err := UnmarshalStorage(data) // transforming []byte into []task.Task
	if err != nil {
		log.Fatal(err)
	}

	if args[1] == "add" {
		id, err := GetId(tasks)
		if err != nil {
			log.Fatal(err)
		}

		err = AddTask(args, tasks, id)
		if err != nil {
			log.Fatal(err, " error during adding task")
		}
		fmt.Printf("Task added successfully (ID: %d)", id)
	}

	if args[1] == "update" {
		findId, err := strconv.Atoi(args[2])
		updatedDesc := args[3]
		if err != nil {
			log.Fatal(err)
		}

		err = UpdateTask(findId, updatedDesc, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	if args[1] == "delete" {
		findId, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal()
		}

		err = DeleteTask(findId, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	if args[1] == "mark-in-progress" {
		findId, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}

		err = MarkInProgress(findId, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	if args[1] == "mark-done" {
		findId, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}

		err = MarkDone(findId, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}

	if args[1] == "list" && len(args) <= 2 {
		ShowList(&tasks)
	} else if args[2] == "done" {
		ShowDone(&tasks)
	} else if args[2] == "todo" {
		ShowToDo(&tasks)
	} else if args[2] == "in-progress" {
		ShowInProgress(&tasks)
	}
}

/*
name - description

task-cli:
add name
update N name
delete N name

mark-in-progress N
mark-done N

list
list done
list todo
list in-progress

*/
