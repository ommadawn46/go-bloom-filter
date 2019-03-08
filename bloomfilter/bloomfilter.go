package bloomfilter

import (
	"crypto/md5"
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

func getDigests(k uint, elem []byte) []uint {
	digest := md5.Sum(elem)
	digests := []uint{}
	n := uint(0)
	for i := 0; i+8 <= len(digest); i++ {
		d1 := uint(binary.BigEndian.Uint64(digest[i : i+8]))
		for j := i + 1; j+8 <= len(digest); j++ {
			d2 := uint(binary.LittleEndian.Uint64(digest[j : j+8]))
			digests, n = append(digests, d1^d2), n+1
			if (1<<n)-1 >= k {
				return digests
			}
		}
	}
	return digests
}

func getHashNum(idx uint, digests []uint) uint {
	mask := idx + 1
	hashNum := uint(0)
	for i := uint(0); (mask >> i) != 0; i++ {
		if (mask>>i)&1 == 0 {
			continue
		}
		hashNum ^= digests[i]
	}
	return hashNum
}

// Add append an element into filter
func (bf *BloomFilter) Add(elem []byte) {
	digests := getDigests(bf.K, elem)
	for i := uint(0); i < bf.K; i++ {
		hashNum := getHashNum(i, digests) % bf.M
		bf.filter[hashNum/8] |= 1 << uint(hashNum%8)
	}
	bf.N++
}

// Contains check an element is included in filter
func (bf *BloomFilter) Contains(elem []byte) bool {
	digests := getDigests(bf.K, elem)
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
