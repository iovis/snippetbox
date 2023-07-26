set dotenv-load

default: list

full:
    tmux new-window -n backend 'just dev'
    tmux new-window -n frontend 'just worker'

# lists available tasks
@list:
    just --list

# start the server
dev:
    go run .

# open the project in the browser
open:
    open "$PROJECT_URL" -a "Google Chrome Canary"

# start a console
# console:
#     rails console

# run tests
test:
    go test

# Open the DB
db:
    pgcli $DATABASE_URL
