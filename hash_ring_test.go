package hashring

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_searchSegments(t *testing.T) {
	type args struct {
		arr []uint64
		val uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"search 2",
			args{[]uint64{4, 8, 15, 16, 23, 42}, 2},
			4,
		},

		{
			"search 4",
			args{[]uint64{4, 8, 15, 16, 23, 42}, 4},
			4,
		},

		{
			"search 9",
			args{[]uint64{4, 8, 15, 16, 23, 42}, 9},
			15,
		},

		{
			"search 43",
			args{[]uint64{4, 8, 15, 16, 23, 42}, 43},
			4,
		},

		{
			"search in empty slice",
			args{[]uint64{}, 2},
			0,
		},

		{
			"search a number in range (0...n]",
			args{[]uint64{2}, 1},
			2,
		},

		{
			"search a number out of range (0...n]",
			args{[]uint64{2}, 33},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRanges(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("searchRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashRing(t *testing.T) {
	t.Run("add shard with defined number of virtual virtualShards", func(t *testing.T) {
		const (
			Replicas = 5
			Shard    = "node-1"
		)
		ring := New(Replicas)
		assert.Nil(t, ring.Add(Shard))

		assert.Len(t, ring.hashes, 5)
		assert.Len(t, ring.shards, 5)

		for _, s := range ring.shards {
			assert.Equal(t, Shard, s)
		}
	})

	t.Run("get a shard by a key", func(t *testing.T) {
		const (
			ShardA   = "shard-a"
			ShardB   = "shard-b"
			Replicas = 100
		)

		ring := New(Replicas)

		assert.Nil(t, ring.Add(ShardA))
		assert.Nil(t, ring.Add(ShardB))

		got := make(map[string]int, 1000)
		for i := 0; i < 1000; i++ {
			shard, err := ring.Get(strconv.Itoa(i))
			assert.Nil(t, err)
			got[shard]++
		}

		assert.Equal(t, 700, got[ShardA])
		assert.Equal(t, 300, got[ShardB])
	})
}
