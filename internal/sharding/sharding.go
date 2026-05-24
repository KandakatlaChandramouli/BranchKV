package sharding

import (
	"hash/fnv"
)

type Sharder struct{}

func NewSharder() *Sharder {
	return &Sharder{}
}

func (s *Sharder) Shard(
	key string,
	total int,
) int {

	h := fnv.New32a()

	h.Write([]byte(key))

	return int(
		h.Sum32(),
	) % total
}
