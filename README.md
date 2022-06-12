# ✨ Go Checkerr ✨
A very small utility package for reducing redundancy when dealing with error handling in Go.

> Example of redundancy:

```go
    body, err := iotuil.ReadFile("file.txt")
    if err != nil {
        log.Fatalf(err)
    }
 ```
 
 > Example of package implementation:
```go
  body, err := iotuil.ReadFile("file.txt")
  checkPanic(err)
 ```
