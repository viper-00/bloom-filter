package bloom

import (
	"hash"
	"hash/fnv"
	"math"
)

// CBF: counting Bloom Filter
type CBF struct {
	m      int // the number of bits used in storage.
	k      int // the number of hashing functions on elements of the set.
	h      hash.Hash32
	bfList []int
}

// NewCountingBloomFilter: Create a counting bloom filter with assigned expect element count and false detect rate.
func NewCountingBloomFilter(totalNumber uint32, falseDetectRate float64) *CBF {
	b := &CBF{h: fnv.New32()}
	b.estimateMK(totalNumber, falseDetectRate)
	b.bfList = make([]int, b.m)
	return b
}

func (b *CBF) estimateMK(number uint32, posibility float64) {
	// m = -1 * (n * lnp) / (ln2)^2
	nFloat := float64(number)
	ln2 := math.Log(2)
	b.m = int(-1 * (nFloat * math.Log(posibility)) / math.Pow(ln2, 2))

	// k = m/n * ln2
	b.k = int(math.Ceil(float64(b.m) / nFloat * ln2))
}

// Add: add element into this cbf structure.
func (b *CBF) Add(element []byte) {
	b.bfSet(element)
}

// Test: test element if exist in sbf structure.
func (b *CBF) Test(element []byte) bool {
	return b.bfTest(element)
}

// Remove: will remove item from this structure.
func (b *CBF) Remove(element []byte) {
	if !b.bfTest(element) {
		return
	}

	for i := 0; i < b.k; i++ {
		listIndex := b.hashFuns(i, element)
		if b.bfList[listIndex] != 0 {
			b.bfList[listIndex]--
		}
	}
}

func (b *CBF) bfSet(data []byte) {
	for i := 0; i < b.k; i++ {
		listIndex := b.hashFuns(i, data)
		b.bfList[listIndex]++
	}
}

func (b *CBF) hashFuns(indexFn int, data []byte) int {
	// Hash function
	b.h.Reset()
	b.h.Write(data)
	hashData := b.h.Sum32()
	hasInt := int(hashData)
	return (hasInt + indexFn) % b.m
}

func (b *CBF) bfTest(data []byte) bool {
	for i := 0; i < b.k; i++ {
		listIndex := b.hashFuns(i, data)
		if b.bfList[listIndex] == 0 {
			return false
		}
	}
	return true
}
