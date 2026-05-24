package checkpoint

import (
	"encoding/json"
	"os"
)

type Snapshot struct {
	State map[string]string
}

func Save(
	path string,
	s Snapshot,
) error {

	data, err := json.Marshal(
		s,
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		path,
		data,
		0644,
	)
}

func Load(
	path string,
) (
	Snapshot,
	error,
) {

	data, err := os.ReadFile(
		path,
	)

	if err != nil {
		return Snapshot{}, err
	}

	var s Snapshot

	err = json.Unmarshal(
		data,
		&s,
	)

	return s, err
}
