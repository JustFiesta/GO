# GO Error handling

Go does not have try-catch-finally statement. Instead it takes approach of multi returns (function can return value and error if needed).

Go also provides type error (which is an interface - so every type that has method Error() is also type error)

(type error ¶)[https://pkg.go.dev/builtin#error]
(Why does Go not have exceptions?)[https://go.dev/doc/faq#exceptions]
(Error handling and Go)[https://go.dev/blog/error-handling-and-go]