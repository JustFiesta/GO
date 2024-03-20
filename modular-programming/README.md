# Modular programming in GO

This folder contains initial excercises to get familiar with go modules and 3rd party code.

<hr>

## Notes

Modular programming is a paradiam for splitting code into smaller chunks (modules). This is made o avoid repetition and help understand stack trace better.

Cause:

1. Do not repeat code (if not needed for clarity)
2. Being able to use 3rd party packages

<hr>

**Tools:**

Get 3rd party dependency/package

```go
go get
```

Install/remove missing dependencies in current working go directory

```go
go mod tidy
```

<hr>

### Package visability

It depends on how we name stuff in code.

Uppercase word - exported/visible (ex. Printf())
Lowercase word - private/not visible (ex. function_name())

<hr>

### Useful links

* [GO specyfication](https://go.dev/ref/spec)

