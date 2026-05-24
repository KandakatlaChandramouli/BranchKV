package state_sync

import (
	"sync/atomic"
)

type Sync struct {
	rounds atomic.Uint64
}

func NewSync() *Sync {

	return &Sync{}
}

func (s *Sync) Run() {
	s.rounds.Add(1)
}

func (s *Sync) Count() uint64 {
	return s.rounds.Load()
}
