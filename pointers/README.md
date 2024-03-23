# Pointers in GO

This folder contains folder excercises and notes on pointers in GO.

<hr>

## Notes

Values in GO are immutable. If one has to do something with value to change it - one need to copy value.

Pointers on the other hand are mutable.

* & - shows address of value
* \* - dereferences the memory address (peeks at the value the address is storing)
* \*int - points to specyfic address location of an specified type - needs dereferencing by * to get value

<hr>

### Important topics

* value semantics - Copied values (copied values - "pass by value")
* pointer semantics - Shared memory (copied values - "pass by value")
* pointer/value semantics heuristics
* escape analysis - go run -gcflags -m (check help)
* values are stored in stack (cheap, fast, works on scopes, after return func stack is freed)
* pointers are stored in heap (more expensive, needs garbage collection, allows to share values in different functions)
* exporting methods!
    "method set" is the set of methods attached to a type. This concept is key to the
    Go's interface mechanism, and it is associated with both the value types and pointer types.
    The method set of a type T consists of all methods with receiver type T.  
        -> These methods can be called using variables of type T.  
    
    The method set of a type *T consists of all methods with receiver *T or T
        -> These methods can be called using variables of type *T.  
        -> It can call methods of the corresponding non-pointer type as well  
    The idea of the method set is integral to how interfaces are implemented and used in Go
