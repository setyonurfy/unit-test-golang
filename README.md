# Go Unit Test

This project demonstrates a simple Go application with unit tests.

## Project Structure
go.mod helper/ hello_world.go hello_world_test.go README.md


## Description

- `helper/hello_world.go`: Contains the implementation of the `HelloWorld` and `HelloTayo` functions.
- `helper/hello_world_test.go`: Contains unit tests for the `HelloWorld` and `HelloTayo` functions.

## Functions

### HelloWorld
```go
func HelloWorld(name string) string

Returns a greeting message by appending "hello" to the given name.

HelloTayo

Running Tests
To run the tests, use the following command:

go test ./...

Requirements
Go 1.21.5 or higher
