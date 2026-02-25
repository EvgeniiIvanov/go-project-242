### Hexlet tests and linter status:
[![Actions Status](https://github.com/EvgeniiIvanov/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/EvgeniiIvanov/go-project-242/actions)

## Description

A command-line tool that calculates the size of files and directories, similar to `du`. Supports recursive traversal, human-readable output, and hidden file inclusion.

## Usage

### How to install

```bash
make build
sudo cp bin/hexlet-path-size /usr/local/bin/
```

### How to use

```bash
hexlet-path-size [options] <path>
```

Options:
- `-r, --recursive` - calculate size recursively for directories
- `-H, --human` - display size in human-readable format (KB, MB, GB)
- `-a, --all` - include hidden files and directories

Examples:
```bash
hexlet-path-size .
hexlet-path-size -r -H /var/log
hexlet-path-size -r -a -H ~/Documents
```

### How to uninstall

```bash
sudo rm /usr/local/bin/hexlet-path-size
```

## Asciinema demo

[![asciicast](https://asciinema.org/a/797330.svg)](https://asciinema.org/a/797330)

## Development part

### Project structure

```bash
.
├── Makefile
├── cmd
│   └── hexlet-path-size
│       └── main.go
├── go.mod
├── go.sum
├── path_size.go
├── path_size_test.go
└── testdata
    ├── .hidden
    │   ├── .hidden.txt
    │   └── file.txt
    ├── empty
    |── empty.txt
    ├── goodbye.link
    ├── hello.txt
    ├── subdir
    │   ├── empty
    │   ├── goodby.txt
    │   └── whatsup.txt
    └── .hidden.txt
```

### How to lint

Perform `make lint`

### How to run tests

Perform `make test`

### How to build

Perform `make build`
