package checkpoint

import (
	"encoding/gob"
	"os"
)

type Checkpoint struct {
	Path string
}

func NewCheckpoint(
	path string,
) *Checkpoint {

	return &Checkpoint{
		Path: path,
	}
}

func (c *Checkpoint) Save(
	v any,
) error {

	f, err := os.Create(c.Path)
	if err != nil {
		return err
	}

	defer f.Close()

	return gob.NewEncoder(f).Encode(v)
}
