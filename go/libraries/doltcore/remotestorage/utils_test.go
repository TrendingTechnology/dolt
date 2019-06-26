package remotestorage

import (
	"github.com/liquidata-inc/ld/dolt/go/store/hash"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestHashesToSlices(t *testing.T) {
	const numHashes = 32

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var randomHashes []hash.Hash
	var randomHashBytes [][]byte
	for i := 0; i < numHashes; i++ {
		var h hash.Hash

		for j := 0; j < len(h); j++ {
			h[j] = byte(rng.Intn(255))
		}

		randomHashes = append(randomHashes, h)
		randomHashBytes = append(randomHashBytes, h[:])
	}

	var zeroHash hash.Hash
	tests := []struct {
		name     string
		in       []hash.Hash
		expected [][]byte
	}{
		{
			"test nil",
			nil,
			[][]byte{},
		},
		{
			"test empty",
			[]hash.Hash{},
			[][]byte{},
		},
		{
			"test one hash",
			[]hash.Hash{zeroHash},
			[][]byte{zeroHash[:]},
		},
		{
			"test many random hashes",
			randomHashes,
			randomHashBytes,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := HashesToSlices(test.in)

			if !reflect.DeepEqual(test.expected, actual) {
				t.Error("unexpected result")
			}
		})
	}
}

func TestHashSetToSlices(t *testing.T) {
	const numHashes = 32

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomHashSet := make(hash.HashSet)

	var randomHashes []hash.Hash
	var randomHashBytes [][]byte
	for i := 0; i < numHashes; i++ {
		var h hash.Hash

		for j := 0; j < len(h); j++ {
			h[j] = byte(rng.Intn(255))
		}

		randomHashSet.Insert(h)
		randomHashes = append(randomHashes, h)
		randomHashBytes = append(randomHashBytes, h[:])
	}

	var zeroHash hash.Hash
	tests := []struct {
		name           string
		hashes         hash.HashSet
		expectedHashes []hash.Hash
		expectedBytes  [][]byte
	}{
		{
			"test nil",
			nil,
			[]hash.Hash{},
			[][]byte{},
		},
		{
			"test empty",
			hash.HashSet{},
			[]hash.Hash{},
			[][]byte{},
		},
		{
			"test one hash",
			hash.HashSet{zeroHash: struct{}{}},
			[]hash.Hash{zeroHash},
			[][]byte{zeroHash[:]},
		},
		{
			"test many random hashes",
			randomHashSet,
			randomHashes,
			randomHashBytes,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hashes, bytes := HashSetToSlices(test.hashes)

			if len(hashes) != len(test.hashes) || len(bytes) != len(test.hashes) {
				t.Error("unexpected size")
			}

			for i := 0; i < len(test.hashes); i++ {
				h, hBytes := hashes[i], bytes[i]

				if !test.hashes.Has(h) {
					t.Error("missing hash")
				}

				if !reflect.DeepEqual(h[:], hBytes) {
					t.Error("unexpected bytes")
				}
			}
		})
	}
}

func TestParseByteSlices(t *testing.T) {
	const numHashes = 32

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var randomHashBytes [][]byte
	for i := 0; i < numHashes; i++ {
		var h hash.Hash

		for j := 0; j < len(h); j++ {
			h[j] = byte(rng.Intn(255))
		}

		randomHashBytes = append(randomHashBytes, h[:])
	}

	var zeroHash hash.Hash
	tests := []struct {
		name  string
		bytes [][]byte
	}{
		{
			"test nil",
			[][]byte{},
		},
		{
			"test empty",
			[][]byte{},
		},
		{
			"test one hash",
			[][]byte{zeroHash[:]},
		},
		{
			"test many random hashes",
			randomHashBytes,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hashes, hashToIndex := ParseByteSlices(test.bytes)

			if len(hashes) != len(test.bytes) {
				t.Error("unexpected size")
			}

			for h := range hashes {
				idx := hashToIndex[h]

				if !reflect.DeepEqual(test.bytes[idx], h[:]) {
					t.Error("unexpected value")
				}
			}
		})
	}

}