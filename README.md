# Effective-Mobile Junior Golang task


## Installation
### Locally


Build cli tool:
This project contains my version of the solution on Effective-Mobile task on the Junior/Junior+
Golang Developper

## Installation

### Locally

Build cli tool:
```sh
source config/env.local.prod

make tidy && make build

./

```

You can checkout description of the project and for each command (and their params) in the CLI
```sh
Usage:
  effective-mobile-task [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  migrate     Migrate schema in database
  server      Start server

Flags:
      --dev    use development version
  -h, --help   help for effective-mobile-task
```

Make migrations to the database:

```sh
./effective-mobile-task migrate
```

Start server

```sh
./effective-mobile-task server
```

### In Docker

```sh
source config/env.docker.prod

docker compose -f infra/docker-compose.yaml up -d
```

## Troubleshooting

### Failed to start app service in docker-compose

Try again run docker compose up command and everything should works fine. Really don't how to solve
this problem now, but in a short time i guess patch will be uploaded.
