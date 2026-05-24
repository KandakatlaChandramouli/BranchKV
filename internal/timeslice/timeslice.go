package timeslice

import (
	"time"
)

type Slice struct {
	Duration time.Duration
}

func NewSlice(
	d time.Duration,
) *Slice {

	return &Slice{
		Duration: d,
	}
}
