package hashring

import (
	"fmt"
	"hash/fnv"
)

// FNVHash implements a hash function using hash/fnv algorithm.
func FNVHash(val []byte) (uint64, error) {
	hasher := fnv.New64a()

	if _, err := hasher.Write(val); err != nil {
		return 0, fmt.Errorf("failed to write data to hasher: %w", err)
	}

	return hasher.Sum64(), nil
}
