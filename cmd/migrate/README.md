# Calliope CLI

## Installation

```bash
$ go install lilac.ooo/migrate/cmd/migrate@v1
```

## Usage

```bash
$ migrate -help
```

### Environment variables

The CLI doesn't support environment variables directly, but you can use the usual syntax supported by your shell to pass them as arguments.

```bash
$ migrate -database "$DATABASE_URL"
```
