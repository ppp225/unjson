# unjson
Get JSON values just like in node, without structs, just path with dot notation + brackets

Uses [GJson](https://github.com/tidwall/gjson) internally. Wraps a very limited subset of it to support bracket indexing.

# example

```go
data := unjson.LoadFile("lighthouse.json")
v := unjson.Get(data, "audits.screenshot-thumbnails.details.items[3].timing")
```

# benchmarks

```
Benchmark_LargeFile_unjson-12    	    1236	    818774 ns/op	    4224 B/op	      58 allocs/op
Benchmark_LargeFile_GJson-12     	    1377	    814089 ns/op	    4224 B/op	      58 allocs/op
Benchmark_SmallFile_unjson-12    	  537495	      2141 ns/op	      40 B/op	       3 allocs/op
Benchmark_SmallFile_GJson-12     	  600687	      2054 ns/op	      40 B/op	       3 allocs/op
```