package checksum

import (
	"hash/crc32"
)

type Checksum struct{}

func NewChecksum() *Checksum {
	return &Checksum{}
}

func (c *Checksum) Compute(
	data []byte,
) uint32 {

	return crc32.ChecksumIEEE(data)
}
