# 📝 Task Tracker CLI (Go)

Консольный менеджер задач на Go. Поддерживает добавление, обновление, удаление и изменение статуса задач.  
Данные сохраняются локально в `storage.json`.

---

## 🚀 Установка и запуск

```
cd cmd
go build -o app.exe
.app.exe
```

## 📦 Команды

| Команда | Описание |
|--------|----------|
| add "text" | добавить задачу |
| update <id> "text" | обновить текст задачи |
| delete <id> | удалить задачу |
| mark-in-progress <id> | статус → in-progress |
| mark-done <id> | статус → done |
| list | вывести все задачи |
| list todo | вывести только TODO |
| list done | вывести выполненные |
| list in-progress | вывести задачи в процессе |

---

## Примеры использования

```
.app.exe add "Buy groceries"
# Task added successfully (ID: 1)

.app.exe update 1 "Buy groceries and cook dinner"
.app.exe delete 1

.app.exe mark-in-progress 1
.app.exe mark-done 1

.app.exe list
.app.exe list done
.app.exe list todo
.app.exe list in-progress
```
