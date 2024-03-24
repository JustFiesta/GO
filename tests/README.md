# Tests in GO

This folder contains information about testing functions in go language.

There is a builtin package for testing in go, named: "testing"

<hr>

## Naming conventions

* files
General guideline is to place the test files in the module that they are needed - so if testing main/main.go you need to create a main/main_test.go file.

* functions
Testing functions have own naming convention - func TestXxx(t *testing.T), where Xxx is name of function to test, with capital letter at start

* use t.Error to singal test failure + use "Expected: " ... "got: " idiom

NOTE: remember to BET on your code to run! And to be sure its gonna run:

* Benchmark - benchmark code
* Example - write examples for code in tests
* Test - test code, use table tests (how-test-should-look/ex3)

### Running tests

```shell
go test
```

#### Usefull commands

Check for poor coding style - returns just a suggestions

[Golint](https://github.com/golang/lint)
```shell
golint
```

Formats go code

```shell
gofmt
```

Reposts suspicious constructs

```shell
go vet
```

Benchmark all codesamples:

```shell
go test -bench . 
```

Check test coverage on codebase:

```shell
go test -cover
```

Save coverage data to file:

```shell
go test -coverprofile c.out
```

Run coverage file as a site:

```shell
go tool cover -html=c.out
```
