# go-bloom-filter

```
❯ go run main.go
N: 300000
M: 7188794bits (877KiB)
K: 17
FN: 0.00000000
FP Predict: 0.00001002
FP Actual: 0.00001333
```

```
❯ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/ommadawn46/go-bloom-filter/bloomfilter
BenchmarkAdd-4                 2         892253683 ns/op        13848856 B/op    1700009 allocs/op
BenchmarkContains-4           10         223782231 ns/op         3481412 B/op     435089 allocs/op
PASS
ok      github.com/ommadawn46/go-bloom-filter/bloomfilter       6.053s
```