📝 CLI Task Tracker

Простой и удобный консольный To-Do менеджер, написанный на Go.
Позволяет добавлять задачи, отмечать их выполненными, обновлять и удалять.
Все данные хранятся локально в storage.json.

🚀 Установка и запуск

Клонируйте репозиторий:

git clone https://github.com/venwex/cli-task-tracker.git
cd cli-task-tracker


Соберите приложение:

cd cmd
go build -o app.exe


Запустите:

.\app.exe

📚 Команды
Команда	Описание
add "text"	Добавить задачу
update N "text"	Обновить задачу с ID N
delete N	Удалить задачу
mark-in-progress N	Перевести задачу в статус in-progress
mark-done N	Отметить как done
list	Показать все задачи
list todo	Только todo
list done	Только done
list in-progress	Только in-progress
💡 Примеры использования
.\app.exe add "Buy groceries"
.\app.exe update 1 "Buy groceries and cook dinner"
.\app.exe delete 1

.\app.exe mark-in-progress 1
.\app.exe mark-done 1

.\app.exe list
.\app.exe list todo
.\app.exe list done
.\app.exe list in-progress

📁 Структура проекта
cli-task-tracker/
├── cmd/
│   └── main.go
├── service/
│   └── service.go
├── entities/
│   └── task.go
├── storage/
│   └── storage.json
└── go.mod
