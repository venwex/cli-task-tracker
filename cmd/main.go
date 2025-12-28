package main

import (
	"example/cli-taskTracker/cli"
	"example/cli-taskTracker/service"
)

func main() {
	cli.GetInput()
	service.Router()
}

/*
--- Как запустить проект?

1) перейдите в папку cmd
2) забилдите приложение: go build -o app.exe
3) запускайте программу через: .\\app.exe 

--- список доступных комманд:
add - добавить задачу
update N - обновить N-ую задачу
delete N - удалить N-ую задачу

mark-in-progress N - отметить в прогрессе N-ую задачу
mark-done N - отметить выполнено N-ую задачу

list - вывести все задачи
list done - вывести выполенные задачи
list todo - вывести todo задачи
list in-progress - вывести задачи в прогрессе

--- примеры команд:
.\\app.exe add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
.\\app.exe update 1 "Buy groceries and cook dinner"
.\\app.exe delete 1

# Marking a task as in progress or done
.\\app.exe mark-in-progress 1
.\\app.exe mark-done 1

# Listing all tasks
.\\app.exe list

# Listing tasks by status
.\\app.exe list done
.\\app.exe list todo
.\\app.exe list in-progress
*/
