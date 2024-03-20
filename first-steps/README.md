# First steps at road with GO

This folder contains initial excercises to get familiar with go syntax and workflow.

## Notes

Initilize new project:

```go
go mod init someName
```

Run .go file:

```go
go run filename
```

Build exectuable

```go
go build filename
```

Build for certain OS:

```go
GOOS=darvin go build
GOOS=linux go build
GOOS=windows go build
```

Installing built app:

```go
go install
```

-> this installs app binary in GOPATH/bin

### Useful links

* [GO specyfication](https://go.dev/ref/spec)
* [Standard Library](https://pkg.go.dev/std)
* [Documentation](https://go.dev/doc/)
* [How to write go code](https://go.dev/doc/code)
* [Common mistakes](https://golang50shad.es/)
* [Managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)
* [How to use GO modules](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)
* [How ot create GO modules](https://go.dev/doc/tutorial/create-module)
* [Proverbs](https://go-proverbs.github.io/)
