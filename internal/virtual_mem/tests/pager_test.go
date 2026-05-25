package tests

import (
	"branchkv-core/internal/virtual_mem"
	"branchkv-core/internal/virtual_mem/eviction"
	"branchkv-core/internal/virtual_mem/pager"
	"branchkv-core/internal/virtual_mem/swap"
	"testing"
)

func TestSwapPager(
	t *testing.T,
) {

	swapper := swap.NewManager()

	evictor := eviction.NewEvictor(
		swapper,
	)

	p := pager.NewPager(
		swapper,
	)

	frame := virtual_mem.NewPhysicalFrame(
		1,
		64,
	)

	frame.Data[0] = 99

	frame.MarkDirty()

	evictor.Evict(
		frame,
	)

	frame.Data[0] = 0

	ok := p.Fault(
		frame,
	)

	if !ok {
		t.Fatal("swapin failed")
	}

	if frame.Data[0] != 99 {
		t.Fatal("data corrupted")
	}
}
