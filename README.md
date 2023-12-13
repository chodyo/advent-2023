# Advent of Code 2023: Go Template

[![Go](https://github.com/chodyo/advent-2023/actions/workflows/go.yml/badge.svg)](https://github.com/chodyo/advent-2023/actions/workflows/go.yml)

## Setup Instructions

1. Create a new repo from this template.
2. Implement your algorithms for each day in `internal/commands/<dayXX>.go`.
   1. Each day has 2 parts, you will replace the code in functions named `dayXXpartY`.
3. Test your implementation using `go test -v ./...`

## Running

Using `go run . -h` will print help text that will guide you on how to run this code.

## Testing

`go test -v ./...`
