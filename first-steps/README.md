# First steps at road with GO

This folder contains initial excercises to get familiar with go syntax and workflow.

<hr>

## Notes

Create go workspace (module):

```shell
go mod init someModuleName
```

-> also to even make/use go files in certain folder we need to specify this module name as shown above

Run .go file:

```shell
go run filename
```

Build exectuable

```shell
go build filename
```

Build for certain OS:

```shell
GOOS=darvin go build
GOOS=linux go build
GOOS=windows go build
```

Installing built app:

```shell
go install
```

-> this installs app binary in GOPATH/bin

Checking the sha sum on downlaod file:

```shell
shasum --help
shasum -a path/to/file
```

-> unchanged file shloud have always the same sum as the online version

<hr>

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
