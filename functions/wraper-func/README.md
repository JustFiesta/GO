# Wraper fucntions

This folder contains notes about wraper functions.

<hr>

## Notes

Wraper function is a function that takes as a parameter other function to envoke it and not use its return values.

```go
func(fn func(num int) float){
    fmt.Println("Excecution start")

    fn()

    fmt.Println("Excecution completed")
}
```

```go
func TimeFunction(fn func()){
    start := time.Now()

    fn()

    elapsed := time.Since(start)
    fmt.Println("Elapsed time: ", elapsed)
}
```

<hr>

### Useage

In opposition ot callback - wrapper func can provide additional behavior, or modify it, without direclty modifying wrapped function code.

**It takes same arguments and returns same type as wrapped function** and maight preform pre or post processing, error handling, logging, or other tasks around the execution of wrapped function.
