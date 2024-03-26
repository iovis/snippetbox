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
