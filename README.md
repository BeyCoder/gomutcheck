# gomutcheck
A golangci-lint linter to detect struct field mutations in value receiver methods.

# Example
```go
package gomutcheck

type ExampleStruct struct {
    ExampleField string
}

func (s ExampleStruct) MutateField() {
    s.ExampleField = "new value"
}
```

### Result:
```
example.go:8:2: struct field 'ExampleField' is being mutated in value receiver method
```