# static
[![GoDoc](https://godoc.org/github.com/prologic/static?status.svg)](https://godoc.org/github.com/prologic/static)
[![Go Report Card](https://goreportcard.com/badge/github.com/prologic/static)](https://goreportcard.com/report/github.com/prologic/static)

static is a self-hosted static file serving app which does nothing more than
just serve up static files with a configured root path.

## Installation

### Source

```#!bash
$ go install github.com/prologic/static/...
```

## Usage

Run static:

```#!bash
$ static
```

Then visit: http://localhost:8000/

## Configuration

By default static services up files from the current directory. This can be
configured via the `-root` option.

## License

MIT
