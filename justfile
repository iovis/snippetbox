set dotenv-load := true

bin := "./cmd/web"

default: run

@list:
    just --list

run:
    go run {{ bin }}

build:
    go build {{ bin }}

dev:
    watchexec -re go,html just run

test:
    go test

db:
    pgcli $DATABASE_URL
