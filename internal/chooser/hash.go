package chooser

import (
	"fmt"
	"hash/fnv"
)

func hash(val []byte) (uint64, error) {
	hasher := fnv.New64a()

	if _, err := hasher.Write(val); err != nil {
		return 0, fmt.Errorf("failed to write data to hasher: %w", err)
	}

	return hasher.Sum64(), nil
}