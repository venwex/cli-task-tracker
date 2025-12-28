# 📝 Task Tracker CLI (Go)

Консольный менеджер задач на Go. Поддерживает добавление, обновление, удаление и изменение статуса задач.  
Данные сохраняются локально в `storage.json`.

---

## 🚀 Установка и запуск

```
cd cmd
go build -o app.exe
.pp.exe
```

Linux/macOS:

```
go build -o app && ./app
```

---

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
.pp.exe add "Buy groceries"
# Task added successfully (ID: 1)

.pp.exe update 1 "Buy groceries and cook dinner"
.pp.exe delete 1

.pp.exe mark-in-progress 1
.pp.exe mark-done 1

.pp.exe list
.pp.exe list done
.pp.exe list todo
.pp.exe list in-progress
```

---

## Формат хранения (`storage.json`)

```
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "createdAt": "2025-12-22T00:56:20",
    "updatedAt": "2025-12-22T01:10:42"
  }
]
```
