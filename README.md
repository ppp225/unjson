# unjson
Get JSON values just like in node, without structs, just path

# example

```go
filename := "lighthouse.json"
path := "categories.performance.score"
v := unjson.Getf(filename, path)
fmt.Print(v) // prints performance score
```