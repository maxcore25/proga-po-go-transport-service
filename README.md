# proga-po-go-transport-service

## Инструкция

1. Скачайте зависимости:

```sh
go mod tidy
```

2. Скопируйте `.env.example` и переименуйте в `.env`. Опционально можете изменить значения ENV-переменных

3. Запустите проект:

```sh
# Windows
.\tasks.ps1 dev

# Linux
make dev
```

При запуске бекенда добавляются таблицы и начальные тестовые данные.

4. Опционально можно запустить тесты:

```sh
# Windows
.\tasks.ps1 test

# Linux
make test
```

---

## Start

1. start

```sh
go mod init github.com/user/proga-po-go-transport-service
```

2. Install globally

```sh
# hot reload dev server
go install github.com/air-verse/air@latest

# Swagger
go install github.com/swaggo/swag/cmd/swag@latest

# Pretty output for tests
go install gotest.tools/gotestsum@latest
```

## Commands

### Only for the first time (if app does not work)

```sh
# Init air for hot reload dev server
air init

# Init Swagger API docs
swag init --parseDependency --parseInternal -g ./cmd/app/main.go
```

### Windows

```sh
# if you want to see all commands
.\tasks.ps1 help

# run dev server with hot reaload
.\tasks.ps1 dev
```

### Linux

```sh
make help

make dev
```

## Tests

```sh
# Run all tests (stdlib)
go test ./...

# Run all tests (pretty output)
gotestsum --format testname

# or

# Windows
.\tasks.ps1 test

# Linux
make test
```
