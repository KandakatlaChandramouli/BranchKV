package tests

import (
	"branchkv-core/internal/streaming"
	"testing"
)

func TestStreaming(
	t *testing.T,
) {

	s := streaming.NewStream()

	s.Publish(
		[]byte("runtime"),
	)

	if s.Size() != 1 {
		t.Fatal("stream failed")
	}
}
