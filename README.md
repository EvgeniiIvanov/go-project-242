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

### How to delete

```bash
sudo rm /usr/local/bin/hexlet-path-size
```

## Dev part

### How to lint

Perform `make lint`

### How to run tests

Perform `make test`

### How to build

Perform `make build`

### How to run

```bash
./bin/hexlet-path-size <path>
```