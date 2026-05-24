package main

import (
	"branchkv-core/internal/bloomfilter"
	"branchkv-core/internal/compactionplanner"
	"branchkv-core/internal/lsmtree"
	"branchkv-core/internal/manifest"
	"branchkv-core/internal/memtable"
	"branchkv-core/internal/segmentmanager"
	"branchkv-core/internal/tablecache"
	"branchkv-core/internal/writebuffer"
	"fmt"
)

func main() {

	tree := lsmtree.NewLSMTree()

	tree.Put(
		"runtime",
		[]byte("branchkv"),
	)

	mem := memtable.NewTable()

	mem.Put(
		"hello",
		[]byte("world"),
	)

	buffer := writebuffer.NewBuffer()

	buffer.Append(
		[]byte("wal"),
	)

	filter := bloomfilter.NewFilter(1024)

	filter.Add("runtime")

	planner := compactionplanner.NewPlanner()

	planner.AddPlan(1)

	segments := segmentmanager.NewManager()

	segments.Add(1)

	cache := tablecache.NewCache()

	cache.Add("sst-1")

	manifestFile := manifest.NewManifest()

	manifestFile.Add("manifest-1")

	fmt.Println(
		"tree:",
		tree.Size(),
	)

	fmt.Println(
		"mem:",
		mem.Size(),
	)

	fmt.Println(
		"buffer:",
		buffer.Size(),
	)

	fmt.Println(
		"filter:",
		filter.Contains("runtime"),
	)

	fmt.Println(
		"plans:",
		planner.Size(),
	)

	fmt.Println(
		"segments:",
		segments.Size(),
	)

	fmt.Println(
		"cache:",
		cache.Size(),
	)

	fmt.Println(
		"manifest:",
		manifestFile.Size(),
	)
}
