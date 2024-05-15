# ConMap

[![Go Reference](https://pkg.go.dev/badge/github.com/Winens/conmap)](https://pkg.go.dev/github.com/Winens/conmap)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Winens/conmap/blob/main/LICENSE)

`ConMap` is a Go library that provides a thread-safe map implementation with **type safety**. It is designed to offer safe concurrent access to a map, making it suitable for use in multi-threaded environments.

## Installation

To install `conmap`, use `go get`:

```sh
go get github.com/Winens/conmap
```

## Usage

```go
// Create new thread-safe map.
m := conmap.New[int64, string]()

// Store data without concerning about race condiditon problem.
m.Store(1, "one")
m.Store(2, "two")
m.Store(3, "three")

// Get value from map.
value, ok := m.Load(1)
if !ok {
    println("Key is not exists in the map.")
}

printf("Value: %s\n", value)

// Delete the value from map with the given key.
m.Delete(1)

// Range over map.
m.Range(func(key int, value string) bool {
    printf("Key: %d, Value: %s\n", key, value)
    return true // If you return false, the loop will break.
})

// Get length of keys that in map.
printf("%d keys are in map", m.Len())
```
