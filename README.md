# go-bloom-filter

```
❯ go run main.go
N: 1000000
M: 23962646 (2.86MiB)
K: 17
FN: 0.00000000
FP Predict: 0.00001002
FP Actual: 0.00000800
```

```
❯ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/ommadawn46/go-bloom-filter/bloomfilter
BenchmarkAdd-4                 3         407660192 ns/op        156954210 B/op   2400299 allocs/op
BenchmarkContains-4            5         259917759 ns/op        156036747 B/op   2400199 allocs/op
PASS
ok      github.com/ommadawn46/go-bloom-filter/bloomfilter       4.427s
```
