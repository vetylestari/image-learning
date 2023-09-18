# Renos Go Repository Template

<div align="center">
<img src="logo.png" >
</div>

## Project Architecture

```bash
📦go-starter-template
 ┣ 📂console
 ┃ ┣ 📂command
 ┃ ┃ ┗ 📜command.go
 ┃ ┣ 📜init_schema.go
 ┃ ┗ 📜root.go
 ┣ 📂database
 ┃ ┣ 📂migrations
 ┃ ┃ ┗ 📜.gitignore
 ┃ ┣ 📂seeders
 ┃ ┃ ┗ 📜.gitignore
 ┃ ┗ 📜postgres.go
 ┣ 📂domain
 ┃ ┣ 📜<module_name>.go
 ┃ ┗ 📜README.md
 ┣ 📂infrastructure
 ┃ ┣ 📜env.go
 ┃ ┣ 📜heimdall.go
 ┃ ┣ 📜httpserver.go
 ┃ ┣ 📜logger.go
 ┃ ┣ 📜query_builder.go
 ┃ ┗ 📜validation.go
 ┣ 📂lib
 ┃ ┣ 📂datetime
 ┃ ┃ ┗ 📜datetime.go
 ┃ ┣ 📂generator
 ┃ ┃ ┗ 📜string_generator.go
 ┃ ┣ 📂middleware
 ┃ ┃ ┗ 📜request_validation.go
 ┃ ┣ 📂password
 ┃ ┃ ┗ 📜password.go
 ┃ ┣ 📂response
 ┃ ┃ ┣ 📜check_code.go
 ┃ ┃ ┣ 📜errors.go
 ┃ ┃ ┗ 📜response.go
 ┃ ┣ 📜parse_response.go
 ┃ ┣ 📜README.md
 ┃ ┣ 📜response.go
 ┃ ┗ 📜struct_json.go
 ┣ 📂log
 ┃ ┗ 📜.gitignore
 ┣ 📂<module_name>
 ┃ ┣ 📂delivery
 ┃ ┃ ┗ 📂http
 ┃ ┃ ┃ ┗ 📜<module_name>_handler.go
 ┃ ┣ 📂repositories
 ┃ ┃ ┣ 📂postgres
 ┃ ┃ ┃ ┗ 📜<module_name>_postgres_repo.go
 ┃ ┃ ┗ 📂thirdpartyapi
 ┃ ┃ ┃ ┗ 📜<module_name>_thirdpartyapi_repo.go
 ┃ ┣ 📂usecases
 ┃ ┃ ┗ 📜<module_name>_usecase.go
 ┃ ┗ 📜root.go
 ┣ 📜.env.example
 ┣ 📜.gitignore
 ┣ 📜dev.Dockerfile
 ┣ 📜docker-compose.yaml
 ┣ 📜Dockerfile
 ┣ 📜entrypoint.sh
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜logo.png
 ┣ 📜main.go
 ┣ 📜Makefile
 ┗ 📜README.md
```

Source & Explanation about this architecture:

- [Golang Clean Architecture Example Repo](https://github.com/bxcodec/go-clean-arch)

## Requirements

- Postgres
- CompileDaemon (Hot Reload) (<https://github.com/githubnemo/CompileDaemon>)
- Go (<https://go.dev/dl/>)
- Goose (Go Migration) <https://github.com/pressly/goose#install> <https://pressly.github.io/goose/installation/#linux>
- Docker
- Git
- run makefile (for windows) (<https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058#make>)

## How to use this template

1. Click `use this template` button
2. Replace all occurences of `go-starter-template` to `repo_name` in all files
3. If `repo_name` have dash (-) , change MIGRATION_TABLE_NAME value in env to snake_case
4. Copy .env.example to .env and configure it.
5. Run `bash make init-schema` to create new schema.
6. Edit README.md and remove this section.

## Run for Development (With Docker)

Clone the project

```bash
  git clone https://github.com/Renos-id/go-starter-template.git
```

Go to the project directory

```bash
  cd my-project
```

Copy example env and configure the env

```bash
    cp .env.example .env
```

Run Docker Compose Up

```bash
    docker-compose up --build
```

## Run for Development (With All Requirements installed on your machine)

Clone the project

```bash
  git clone https://github.com/Renos-id/go-starter-template.git
```

Go to the project directory

```bash
  cd my-project
```

Copy example env and configure the env

```bash
    cp .env.example .env
```

Install dependencies

```bash
  go mod download
```

Run Init Schema to create schema

```bash
    make init-schema
```

or

```bash
    go run cmd/command/command.go init-schema
```

Run migrations

```bash
    make migrate-up
```

Run seeders

```bash
    make seed-up
```

Start the server

```bash
  go run main.go
```

Or Run Hot Reload

```bash
  make dev
```

or

```bash
  CompileDaemon -build="go build -a -buildvcs=false -o go-starter-template" -command="./go-starter-template"
```

## Makefile commands

Create New Module

```bash
    make module-new name=module_name
```

Init Schema

```bash
    make init-schema
```

Run Hot Reload for Development

```bash
    make dev
```

Run Go Project

```bash
    make run
```

Build Go Project

```bash
    make build
```

Seed the DB to the most recent version available

```bash
    make seed-up
```

Creates new seeders file with the current timestamp

```bash
    make seed-create name=tests_table_seeders
```

Migrate the DB to the most recent version available

```bash
    make migrate-up
```

Migrate the DB up by 1

```bash
    make migrate-up-by-one
```

Migrate the DB to a specific VERSION

```bash
    make migrate-up-to version=00001
```

Roll back the version by 1

```bash
    make migrate-down
```

Roll back to a specific VERSION

```bash
    make migrate-down-to version=00001
```

Re-run the latest migration

```bash
    make migrate-redo
```

Roll back all migrations

```bash
    make migrate-reset
```

Dump the migration status for the current DB

```bash
    make migrate-status
```

Creates new migration file with the current timestamp

```bash
    make migrate-create name=create_tests_table
```

Apply sequential ordering to migrations

```bash
    make migration-fix
```

## Dependencies / Libraries

- Chi (HTTP Service Router) (<https://github.com/go-chi/chi>)
- Pgx (Postgres Database Driver) (<https://github.com/jackc/pgx>)
- Goqu (Query Builder) (<https://github.com/doug-martin/goqu>)
- GoDotEnv (Read .env files) (<https://github.com/joho/godotenv>)
- Heimdall (HTTP & REST Client) (<https://github.com/gojek/heimdall>)
- Cobra (CLI Command) (<https://github.com/spf13/cobra>)
- Goose (Database Migration & Seeder) (<https://github.com/pressly/goose>)
- null (Handling with Nullable SQL & JSON values) (<https://github.com/guregu/null>)
- go-playground/validator (Validate HTTP Request data) (<https://github.com/go-playground/validator>)
- Logrus (Logging) (<https://github.com/sirupsen/logrus>)
- grpc-go (gRPC Server & Client) (<https://github.com/grpc/grpc-go>)
