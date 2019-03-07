package bloomfilter

import (
	"crypto/sha256"
	"fmt"
	"math"
)

// OptimizeM optimize parameter M
func OptimizeM(p float64, ni uint) uint {
	n := float64(ni)
	return uint(math.Ceil(
		(n * math.Log(p)) / math.Log(
			1.0/math.Pow(2, math.Log(2)),
		),
	))
}

// OptimizeK optimize parameter K
func OptimizeK(mi uint, ni uint) uint {
	m, n := float64(mi), float64(ni)
	return uint(math.Round((m / n) * math.Log(2)))
}

// BloomFilter Class
type BloomFilter struct {
	M, K, N uint
	filter  []uint8
}

// NewBloomFilter construct BloomFilter
func NewBloomFilter(m uint, k uint) *BloomFilter {
	size := m / 8
	if m%8 != 0 {
		size++
	}
	return &BloomFilter{
		M:      uint(m),
		K:      uint(k),
		N:      uint(0),
		filter: make([]byte, size),
	}
}

func (bf *BloomFilter) getHashNum(idx uint, elem []byte) uint {
	digest := sha256.Sum256(append(
		[]byte(fmt.Sprint(idx)),
		elem...,
	))
	hashNum := uint(0)
	for _, b := range digest {
		hashNum = ((hashNum << 8) + uint(b)) % bf.M
	}
	return hashNum
}

// Add append an element into filter
func (bf *BloomFilter) Add(elem []byte) {
	for i := uint(0); i < bf.K; i++ {
		hashNum := bf.getHashNum(i, elem)
		bf.filter[hashNum/8] |= 1 << uint(hashNum%8)
	}
	bf.N++
}

// Contains check an element is included in filter
func (bf *BloomFilter) Contains(elem []byte) bool {
	for i := uint(0); i < bf.K; i++ {
		hashNum := bf.getHashNum(i, elem)
		if bf.filter[hashNum/8]&(1<<uint(hashNum%8)) == 0 {
			return false
		}
	}
	return true
}

// FalsePositiveProbability calc probability of false positive
func (bf *BloomFilter) FalsePositiveProbability() float64 {
	m, k, n := float64(bf.M), float64(bf.K), float64(bf.N)
	return math.Pow(
		1.0-math.Exp(-(k*n)/m),
		k,
	)
}
