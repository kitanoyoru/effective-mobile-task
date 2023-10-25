# Effective-Mobile Junior Golang task

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

You can checkout description of project and each command (and their params) in the CLI
```
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

docker compose up -f infra/docker-compose.yaml -d
```
