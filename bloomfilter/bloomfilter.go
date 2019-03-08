package bloomfilter

import (
	"crypto/sha256"
	"encoding/binary"
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

func getDigests(elem []byte) []uint {
	digest := sha256.Sum256(elem)
	digests := []uint{}
	for i := 0; i < len(digest)-8; i++ {
		digests = append(
			digests,
			uint(binary.BigEndian.Uint64(digest[i:i+8])),
		)
	}
	return digests
}

func getHashNum(idx uint, digests []uint) uint {
	hashNum := uint(0)
	for i := uint(0); i < uint(len(digests)); i++ {
		if ((idx+1)>>i)&1 == 0 {
			continue
		}
		hashNum ^= digests[i]
	}
	return hashNum
}

// Add append an element into filter
func (bf *BloomFilter) Add(elem []byte) {
	digests := getDigests(elem)
	for i := uint(0); i < bf.K; i++ {
		hashNum := getHashNum(i, digests) % bf.M
		bf.filter[hashNum/8] |= 1 << uint(hashNum%8)
	}
	bf.N++
}

// Contains check an element is included in filter
func (bf *BloomFilter) Contains(elem []byte) bool {
	digests := getDigests(elem)
	for i := uint(0); i < bf.K; i++ {
		hashNum := getHashNum(i, digests) % bf.M
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
