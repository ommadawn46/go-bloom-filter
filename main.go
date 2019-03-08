package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ommadawn46/go-bloom-filter/bloomfilter"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p, n := 1e-5, uint(1000000)

	m := bloomfilter.OptimizeM(p, n)
	k := bloomfilter.OptimizeK(m, n)
	bf := bloomfilter.NewBloomFilter(m, k)

	baseStr := fmt.Sprint(rand.Intn(1 << 32))

	// bloomfilterにデータを追加
	for i := uint(0); i < n; i++ {
		bf.Add([]byte(baseStr + fmt.Sprint(i)))
	}

	// bloomfilterに登録したデータで判定し偽陰性率を計算
	fnCount := 0
	for i := uint(0); i < n; i++ {
		if !bf.Contains([]byte(baseStr + fmt.Sprint(i))) {
			fnCount++
		}
	}
	fn := float64(fnCount) / float64(n)

	// bloomfilterに登録していないデータで判定し偽陽性率を計算
	fpCount := 0
	for i := n; i < n*2; i++ {
		if bf.Contains([]byte(baseStr + fmt.Sprint(i))) {
			fpCount++
		}
	}
	fp := float64(fpCount) / float64(n)

	fmt.Printf(
		"N: %d\nM: %d (%.2fMiB)\nK: %d\nFN: %.8f\nFP Predict: %.8f\nFP Actual: %.8f\n",
		bf.N, bf.M, float64(bf.M)/(8*1024*1024), bf.K, fn, bf.FalsePositiveProbability(), fp,
	)
}
