set dotenv-load := true

bin := "./cmd/web"

default: run

@list:
    just --list

templ:
    templ generate

run: templ
    go run {{ bin }}

build: templ
    go build {{ bin }}

dev:
    watchexec -re go,templ just run

open:
    open $PROJECT_URL

test:
    go test

docs:
    tmux new-window -Sn docs 'python3 -m http.server -b 0.0.0.0 -d ~/Downloads/lets-go/html'

db:
    mycli $DATABASE_URL

db_init:
    cat db/migrations/*.sql | mycli $DATABASE_SERVER

db_reset:
    mycli $DATABASE_SERVER -e "drop database $DATABASE_NAME"
    just db_init
