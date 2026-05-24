package ratelimiter

import (
	"sync/atomic"
)

type Limiter struct {
	requests atomic.Uint64
	limit    uint64
}

func NewLimiter(
	limit uint64,
) *Limiter {

	return &Limiter{
		limit: limit,
	}
}

func (l *Limiter) Allow() bool {

	v := l.requests.Add(1)

	return v <= l.limit
}
