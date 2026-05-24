package io_uring

import (
	"sync/atomic"
)

type SubmissionQueue struct {
	entries atomic.Uint64
}

func NewSubmissionQueue() *SubmissionQueue {

	return &SubmissionQueue{}
}

func (s *SubmissionQueue) Submit() {
	s.entries.Add(1)
}

func (s *SubmissionQueue) Count() uint64 {
	return s.entries.Load()
}
