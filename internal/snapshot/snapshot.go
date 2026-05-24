package snapshot

import (
	"encoding/gob"
	"os"
)

type SnapshotEngine struct {
	Path string
}

func NewSnapshotEngine(
	path string,
) *SnapshotEngine {

	return &SnapshotEngine{
		Path: path,
	}
}

func (s *SnapshotEngine) Save(
	v any,
) error {

	f, err := os.Create(s.Path)
	if err != nil {
		return err
	}

	defer f.Close()

	enc := gob.NewEncoder(f)

	return enc.Encode(v)
}

func (s *SnapshotEngine) Load(
	v any,
) error {

	f, err := os.Open(s.Path)
	if err != nil {
		return err
	}

	defer f.Close()

	dec := gob.NewDecoder(f)

	return dec.Decode(v)
}
