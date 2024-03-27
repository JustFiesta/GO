# Creating documentation for GO code

This folder contains info about basic know-how for creation of good documentation in golang.

<hr>

## How to write good GO documentation?

In go everything is about ease of programming. Documentation also. Write comments above code as full sentances to create human readable documentation, explaning what function/struct does.

Good practise is to start comment from the function/package name.

Tests in go are tightly coupled with writing documentation. Check tests/how-test-should-look folder.

**Export** functions/structs if you want to document them. Unexported ones will be omited by tools.

* (Godoc: documenting Go code)[https://go.dev/blog/godoc]
* (Example)[https://go.dev/src/errors/errors.go#L9]

<hr>

### Tools for creating documentation

Preview of package.symbol.mehod documentation

```shell
go doc
```

Tool for creating html website from documentation

```shell
godoc
```

3rd party documentation: **godoc.org**. Your code can be added there also:

* put the (github) url of your code into godoc
* your documentation will appear on godoc
* “refresh” at bottom of page if it is ever out of date

To delete it from godoc.org just remove code from site e.g. github
