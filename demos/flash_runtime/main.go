package main

import (
	"branchkv-core/internal/admissioncontrol"
	"branchkv-core/internal/blocktable"
	"branchkv-core/internal/continuousbatching"
	"branchkv-core/internal/flashattention"
	"branchkv-core/internal/kvpager"
	"branchkv-core/internal/prefixcache"
	"branchkv-core/internal/specdecoder"
	"branchkv-core/internal/tokenallocator"
	"branchkv-core/internal/types"
	"fmt"
)

func main() {

	flash := flashattention.NewFlashAttention()

	out := flash.Compute(
		[]float64{1, 2},
		[]float64{3, 4},
		[]float64{5, 6},
	)

	pager := kvpager.NewKVPager()

	pager.Allocate(
		1,
		256,
	)

	table := blocktable.NewTable()

	table.Insert(1)

	alloc := tokenallocator.NewAllocator()

	alloc.Allocate()

	cache := prefixcache.NewCache()

	cache.Store(
		"runtime",
		[]float32{
			1, 2, 3,
		},
	)

	spec := specdecoder.NewSpecDecoder()

	pred := spec.Predict(
		[]types.TokenID{
			1, 2, 3,
		},
	)

	batch := continuousbatching.NewBatcher()

	batch.Add(1)

	admission := admissioncontrol.NewAdmissionControl()

	admission.Admit(true)

	fmt.Println(
		"flash:",
		len(out),
	)

	fmt.Println(
		"pages:",
		pager.Size(),
	)

	fmt.Println(
		"blocks:",
		table.Size(),
	)

	fmt.Println(
		"tokens:",
		alloc.Count(),
	)

	fmt.Println(
		"cache:",
		cache.Size(),
	)

	fmt.Println(
		"pred:",
		len(pred),
	)

	fmt.Println(
		"batch:",
		batch.Size(),
	)

	fmt.Println(
		"admitted:",
		admission.Admitted(),
	)
}
