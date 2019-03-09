package bloomfilter

import (
	"crypto/md5"
	"encoding/binary"
	"math"
)

// OptimizeM optimize parameter M
func OptimizeM(p float64, n uint) uint {
	c := math.Log(1.0 / math.Pow(2, math.Ln2))
	return uint(math.Ceil((math.Log(p) * float64(n)) / c))
}

// OptimizeK optimize parameter K
func OptimizeK(m uint, n uint) uint {
	return uint(math.Round((float64(m) / float64(n)) * math.Ln2))
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
		M: uint(m), K: uint(k), N: uint(0),
		filter: make([]byte, size),
	}
}

func genDigests(elem []byte, k uint) []uint {
	digest := md5.Sum(elem)
	digests := []uint{}

	n := uint(len(digest) - 8)
	for i := uint(0); (1<<i)-1 < k; i++ {
		j, k := i/n, i%n
		bePart := uint(binary.BigEndian.Uint64(digest[j : j+8]))
		lePart := uint(binary.LittleEndian.Uint64(digest[k : k+8]))
		digests = append(digests, bePart+lePart)
	}
	return digests
}

func genHashNum(digests []uint, idx uint) uint {
	mask := idx + 1
	hashNum := uint(0)
	for i := uint(0); (mask >> i) != 0; i++ {
		if (mask>>i)&1 == 0 {
			continue
		}
		hashNum += digests[i]
	}
	return hashNum
}

// Add append an element into filter
func (bf *BloomFilter) Add(elem []byte) {
	digests := genDigests(elem, bf.K)
	for i := uint(0); i < bf.K; i++ {
		hashNum := genHashNum(digests, i) % bf.M
		bf.filter[hashNum/8] |= 1 << uint(hashNum%8)
	}
	bf.N++
}

// Contains check an element is included in filter
func (bf *BloomFilter) Contains(elem []byte) bool {
	digests := genDigests(elem, bf.K)
	for i := uint(0); i < bf.K; i++ {
		hashNum := genHashNum(digests, i) % bf.M
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
