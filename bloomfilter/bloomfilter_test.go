package bloomfilter

import (
	"fmt"
	"testing"
)

var p float64
var m, k, n uint
var bf *BloomFilter

func init() {
	p, n = 1e-5, uint(300000)
	m = OptimizeM(p, n)
	k = OptimizeK(m, n)
	bf = NewBloomFilter(m, k)
	for i := uint(0); i < n; i++ {
		bf.Add([]byte(fmt.Sprint(i)))
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpBf := NewBloomFilter(m, k)
		baseStr := fmt.Sprint(i)
		for j := uint(0); j < n; j++ {
			tmpBf.Add([]byte(baseStr + fmt.Sprint(j)))
		}
	}
}

func BenchmarkContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		baseStr := fmt.Sprint(i)
		for j := uint(0); j < n; j++ {
			bf.Contains([]byte(baseStr + fmt.Sprint(j)))
		}
	}
}
