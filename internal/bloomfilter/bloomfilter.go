package bloomfilter

import (
	"hash/fnv"
)

type Filter struct {
	bits []bool
}

func NewFilter(
	size int,
) *Filter {

	return &Filter{
		bits: make(
			[]bool,
			size,
		),
	}
}

func (f *Filter) hash(
	key string,
) uint32 {

	h := fnv.New32a()

	h.Write([]byte(key))

	return h.Sum32()
}

func (f *Filter) Add(
	key string,
) {

	idx := f.hash(key) %
		uint32(len(f.bits))

	f.bits[idx] = true
}

func (f *Filter) Contains(
	key string,
) bool {

	idx := f.hash(key) %
		uint32(len(f.bits))

	return f.bits[idx]
}
