# Task Tracker CLI

Небольшой CLI-трекер задач на Go с хранением в JSON-файле.

## Что умеет

- Добавлять задачи
- Обновлять описание задачи
- Переводить задачу в статусы `in-progress` и `done`
- Удалять задачи
- Показывать все задачи или только задачи с конкретным статусом

## Стек

- Go `1.25.6`
- [cobra](https://github.com/spf13/cobra) для CLI
- JSON-файл (`tasks.json`) как хранилище

## Архитектура

Проект разделён по слоям:

- `internal/domain` — доменные сущности, статусы и бизнес-валидация
- `internal/service` — use-case слой (операции над задачами)
- `internal/ports` — контракты репозитория
- `internal/adapters/repository` — JSON-реализация репозитория
- `internal/cli` — команды и вывод в консоль
- `cmd/tracker` — точка входа приложения

Это хорошее разделение ответственности: доменная логика изолирована
от инфраструктуры, что упрощает тестирование и развитие.

## Быстрый старт

### 1. Установка зависимостей

```bash
go mod download
```

### 2. Запуск

```bash
go run ./cmd/tracker
```

или собрать бинарник:

```bash
go build -o task-tracker ./cmd/tracker
./task-tracker
```

По умолчанию задачи сохраняются в файл `tasks.json` в корне проекта.

## Команды

### Добавление задачи

```bash
go run ./cmd/tracker add "Купить молоко"
```

### Обновление описания

```bash
go run ./cmd/tracker update 1 "Купить молоко и хлеб"
```

### Смена статуса

```bash
go run ./cmd/tracker mark-in-progress 1
go run ./cmd/tracker mark-done 1
```

### Удаление

```bash
go run ./cmd/tracker delete 1
```

### Просмотр задач

```bash
go run ./cmd/tracker list
go run ./cmd/tracker list todo
go run ./cmd/tracker list in-progress
go run ./cmd/tracker list done
```

## Тесты

Запуск всех тестов:

```bash
go test ./...
```

На текущий момент есть тесты для:

- `internal/domain/task.go`
- `internal/service/task_service.go`
- `internal/adapters/repository/json_repository.go`

## Формат хранения

Репозиторий хранит данные в JSON-объекте вида:

```json
{
  "last_id": 2,
  "tasks": [
    {
      "ID": 1,
      "Description": "Купить молоко",
      "Status": "todo",
      "CreatedAt": "2026-05-05T12:00:00Z",
      "UpdatedAt": "2026-05-05T12:00:00Z"
    }
  ]
}
```

## Что уже хорошо в проекте

- Чистое слоистое разделение после рефакторинга
- Валидация в доменном слое (`empty description`, `invalid status`)
- Потокобезопасность репозитория через `sync.RWMutex`
- Атомарная запись через временный файл и `os.Rename`
- Базовые автоматические тесты на ключевые сценарии

## Что можно улучшить дальше

- Добавить тесты на `internal/cli` (табличные тесты команд и сообщений)
- Добавить сортировку/пагинацию и фильтры по дате
- Добавить флаги для пути к файлу данных (например, `--db-path`)
- Вынести конфигурацию в `configs/`
- Подготовить release-скрипт или `Makefile`
