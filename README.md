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
BenchmarkAdd-4                 5         246502854 ns/op        41715152 B/op    1800073 allocs/op
BenchmarkContains-4           10         154117491 ns/op        40810011 B/op    1800049 allocs/op
PASS
ok      github.com/ommadawn46/go-bloom-filter/bloomfilter       4.179s
```
