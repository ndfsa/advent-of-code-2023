# Advent of Code 2023

All days are structured as libraries, to run them, run the corresponding test:

```sh
go test ./dayXX
```

Or run all tests:

```sh
go test ./...
```

## antlr4 packages

Some days require antlr4 to parse inputs, to generate all grammars just run:

```sh
go generate ./...
```
