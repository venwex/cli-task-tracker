````md
# CLI Task Tracker

## Как запустить проект

1. Перейдите в папку `cmd`:
   ```bash
   cd cmd
````

2. Скомпилируйте приложение:

   ```bash
   go build -o app.exe
   ```

3. Запустите программу:

   ```bash
   .\app.exe
   ```

---

## Список доступных команд

### Основные операции с задачами

* `add "<description>"` — добавить задачу
  *Пример:* `add "Buy groceries"`

* `update <id> "<new description>"` — обновить задачу с указанным ID
  *Пример:* `update 1 "Buy groceries and cook dinner"`

* `delete <id>` — удалить задачу по ID
  *Пример:* `delete 1`

### Изменение статуса

* `mark-in-progress <id>` — отметить задачу как **in progress**
  *Пример:* `mark-in-progress 1`

* `mark-done <id>` — отметить задачу как **done**
  *Пример:* `mark-done 1`

### Вывод списка задач

* `list` — вывести **все** задачи
* `list done` — вывести **выполненные** задачи
* `list todo` — вывести **todo** задачи
* `list in-progress` — вывести задачи **в прогрессе**

---

## Примеры использования

### Добавление задачи

```bash
.\app.exe add "Buy groceries"
# Output:
# Task added successfully (ID: 1)
```

### Обновление и удаление задачи

```bash
.\app.exe update 1 "Buy groceries and cook dinner"
.\app.exe delete 1
```

### Изменение статуса задачи

```bash
.\app.exe mark-in-progress 1
.\app.exe mark-done 1
```

### Просмотр задач

```bash
# Все задачи
.\app.exe list

# Выполненные задачи
.\app.exe list done

# Задачи в статусе "todo"
.\app.exe list todo

# Задачи в прогрессе
.\app.exe list in-progress
```

```
```
