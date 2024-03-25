# Concurrency in GO

^ means writing code to run in indepandant fashon from each funcion/module/etc (similar to bash command&). This enables code to be ran in parallel (like two cars side by side on road).

## Notes

* main is a goroutine

* goroutines are closed when main ends by default. To make them finish use `sync` package
  * Note: concurrent programming might be unprecitable. Eg. Sometimes goroutines might run and sometimes main will finish first

* to create new goroutine use `go func(){}`, ex.:

```go
func hello(){
    fmt.Println("Hello")
}

go hello()
```

! By default this will run in background (just as bash & command)

To run it in CLI:

```go
for i := 0; i < 5; i++ {
    go hello()
}
```

* to check how many are running (or check different info about system) use `runtime` package

```go
runtime.NumGoroutine()
```

* routines synchronisation - package `sync` does that:

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(2)

    go func(){
        hello()
        wg.Done() 
    }()
}

wg.Wait() 
```

* Sometimes few goruntimes are operating on the same data - this cloud make an race condition.

To check for data runtime races:

```shell
go run -race main.go
```

### Usefull links

* (sync package)[https://pkg.go.dev/sync]
* (runtime package)[https://pkg.go.dev/runtime]
* (Rob Pike on Concurrency and Parallelism)[https://www.youtube.com/watch?v=oV9rvDllKEg&ab_channel=gnbitcom]
* (Go concurrency Patterns)[https://go.dev/blog/pipelines]
* (Data Race Detecrot)[https://go.dev/doc/articles/race_detector]