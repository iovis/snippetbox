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

docs:
    tmux new-window -Sn docs 'python3 -m http.server -b 0.0.0.0 -d ~/Downloads/lets-go/html'

db:
    pgcli $DATABASE_URL
