# Group Manager

**Фамилия Имя:** Куркин Николай

**Тестовое задание на Golang**

REST API сервис для управления иерархическими группами людей с рекурсивным подсчётом участников.

## Описание проекта

Сервис предоставляет:

- Создание, обновление и удаление групп с древовидной иерархией (родительские/дочерние группы)
- Создание, обновление и удаление людей (имя, фамилия, год рождения) — при создании привязка к группе, при обновлении возможность смены группы
- Список всех групп с количеством участников: только в данной группе (`direct_count`) и общее количество вместе с дочерними (`total_count`) через рекурсивный CTE
- Список участников группы — только привязанных к данной группе или со всеми дочерними группами (`?deep=true`)

**Технологии:** Go, PostgreSQL, goose (миграции), sqlc, chi, Snowflake ID, Docker.

**Архитектура:** Чистая архитектура (entity -> usecase -> infra/postgres -> delivery/http), внедрение зависимостей.

## Подготовительные действия

### Требования

- Go 1.26+
- Docker и Docker Compose
- (Опционально) goose - для запуска миграций вручную

### Настройка

1. Склонировать репозиторий:
   ```bash
   git clone <repo-url>
   cd group-manager
   ```

2. Настройка окружения (файл `.env` уже содержит значения по умолчанию):
   ```env
   POSTGRES_DB=group_manager
   POSTGRES_USER=group_manager
   POSTGRES_PASSWORD=supersecret
   ```

3. Конфигурация приложения (`config/config.yaml`):
   ```yaml
   http:
     host: "0.0.0.0"
     port: 50000
   postgres:
     conn_string: "host=localhost port=5432 user=group_manager password=supersecret dbname=group_manager sslmode=disable"
   snowflake:
     node_id: 1
   ```

## Информация о доступах

### База данных (PostgreSQL)

| Параметр | Значение |
|---|---|
| Host | `localhost` |
| Port | `5432` |
| Database | `group_manager` |
| User | `group_manager` |
| Password | `supersecret` |
| Connection string | `postgres://group_manager:supersecret@localhost:5432/group_manager?sslmode=disable` |

### API

| Параметр | Значение |
|---|---|
| Host | `http://localhost` |
| Port | `50000` |
| Base path | `/v1` |

## Как запустить проект

### Быстрый старт (Docker Compose)

```bash
# Запустить PostgreSQL + API + миграции
just up

# Проверить статус
just status

# Посмотреть логи
just logs
```

### Запуск вручную

```bash
# 1. Запустить PostgreSQL
docker compose -f deploy/compose.yaml -p group_manager --env-file .env up -d postgres

# 2. Применить миграции
goose -dir migrations up

# 3. Запустить API
go run ./cmd/api
```

### Остановка

```bash
just down
```

### Доступные команды (Justfile)

| Команда | Описание |
|---|---|
| `just up` | Запустить все сервисы |
| `just down` | Остановить все сервисы |
| `just restart` | Перезапустить сервисы |
| `just logs` | Логи сервисов |
| `just status` | Статус сервисов |
| `just postgres` | Подключиться к psql |
| `just sqlc` | Перегенерировать код sqlc |

### Примеры запросов

```bash
# Создать группу
curl -X POST http://localhost:50000/v1/groups \
  -H "Content-Type: application/json" \
  -d '{"name":"Backend","parent_id":null}'

# Создать человека
curl -X POST http://localhost:50000/v1/people \
  -H "Content-Type: application/json" \
  -d '{"group_id":1,"firstname":"Иван","lastname":"Иванов","birthday":"1990-01-15"}'

# Список групп с количеством участников
curl http://localhost:50000/v1/groups

# Список участников группы (с дочерними)
curl "http://localhost:50000/v1/groups/1/members?deep=true"
```
