# Tests in GO

This folder contains information about testing functions in go language.

There is a builtin package for testing in go, named: "testing"

<hr>

## Naming conventions

* files
General guideline is to place the test files in the module that they are needed - so if testing main/main.go you need to create a main/main_test.go file.

* functions
Testing functions have own naming convention - func TestXxx(t *testing.T), where Xxx is name of function to test, with capital letter at start

### Running tests

```shell
go test
```
