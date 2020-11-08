# Mapoteca

Server responsible for storing (in a database) projects.

## Project dependencies

### Golang

Install [Go version 1.14](https://godoc.org/golang.org/dl/go1.14)

### PostgreSQL

Install PostreSQL through your favorite package manager.

Create your database and make sure to fill the environment variables at `.env` after reading the **Project configurations** topic.

#### Configure your database

Create a database with the name of your choice using the handy `createdb` command:

```sh
$ createdb mapoteca_localhost
```

Access it and create a user of your choice (not need to follow example below):

```sh
$ psql mapoteca_localhost

mapoteca_localhost=# CREATE ROLE mapoteca_user WITH LOGIN PASSWORD 'password';
```

## Project configurations

We use `dotenv` for managing environment variables. Check the `.env.sample` in the root folder to check the documentation for all necessary variables to be configured.

Create a copy of `.env.sample`, name it as `.env` and fill its values according to the documentation inside the file.

## Run server

After following the steps above you may simply do:

```sh
go run main.go
```

