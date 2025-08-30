[![Documentation](https://pkg.go.dev/badge/lilac.ooo/migrate)](https://pkg.go.dev/lilac.ooo/migrate)

# Calliope

A minimalist fork of [golang-migrate](https://github.com/golang-migrate/migrate). It loosens some opinions of its sibling, while tightening others.

- It doesn't require down migrations, or for up migrations to have ".up" in their name when a down migration is not present.
- It only supports reading migrations from the local disk, or an `io.FS` when used as a library.
- It only support Postgres and SQLite 3.
- It only supports files that use .sql extensions.


## Databases

Calliope has a very focused set of supported databases. If the database you want isn't supported, you should probably use golang-migrate instead.

* [PostgreSQL](database/postgres)
* [SQLite3](database/sqlite3)

### Database URLs

Database connection strings are specified via URLs. The URL format is driver dependent but generally has the form: `dbdriver://username:password@host:port/dbname?param1=true&param2=false`

As typical of URLs, they need to be URL encoded. Unencoded characters like `!` or `#` will result in unexpected behavior.

## CLI usage

* Handles ctrl+c (SIGINT) gracefully.
* No config search paths, no config files, no magic ENV var injections.

[CLI Documentation](cmd/migrate) (includes CLI install instructions)

### Basic usage

```bash
$ migrate -path migrations/ -database postgres:///database up
```

## Use in your Go project

* API is stable and frozen for this release (v3 & v4).
* Uses [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies.
* To help prevent database corruptions, it supports graceful stops via `GracefulStop chan bool`.
* Bring your own logger.
* Uses `io.Reader` streams internally for low memory overhead.
* Thread-safe and no goroutine leaks.

__[Go Documentation](https://pkg.go.dev/lilac.ooo/migrate)__

```go
import (
    "lilac.ooo/migrate"
    _ "lilac.ooo/migrate/database/postgres"
)

func main() {
    m, err := migrate.New(
        "file://migrations/",
        "postgres://localhost:5432/database?sslmode=enable")
    m.Steps(2)
}
```

Want to use an existing database client?

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
    "lilac.ooo/migrate"
    "lilac.ooo/migrate/database/postgres"
    _ "lilac.ooo/migrate/source/file"
)

func main() {
    db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file:///migrations",
        "postgres", driver)
    m.Up() // or m.Steps(2) if you want to explicitly set the number of migrations to run
}
```

## Getting started

Go to [getting started](GETTING_STARTED.md)
