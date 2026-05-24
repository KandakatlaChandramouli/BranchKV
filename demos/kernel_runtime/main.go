package main

import (
	"branchkv-core/internal/branchgraph"
	"branchkv-core/internal/epoch"
	"branchkv-core/internal/logstructured"
	"branchkv-core/internal/pagecache"
	"branchkv-core/internal/runtime_gc"
	"branchkv-core/internal/segment"
	"branchkv-core/internal/vectordb"
	"fmt"
)

func main() {

	cache := pagecache.NewPageCache()

	cache.Store(
		&pagecache.CachedPage{
			ID:   1,
			Data: []byte("runtime"),
		},
	)

	db := vectordb.NewVectorDB()

	db.Insert(
		1,
		[]float32{
			1, 2, 3,
		},
	)

	store := logstructured.NewLogStructuredStore()

	store.Append(
		logstructured.Record{
			Key:   "branch",
			Value: []byte("state"),
		},
	)

	graph := branchgraph.NewBranchGraph()

	graph.Connect(1, 2)

	gc := runtime_gc.NewRuntimeGC()

	gc.Collect()

	epochManager := epoch.NewEpochManager()

	epochValue := epochManager.Next()

	segments := segment.NewSegmentManager()

	segments.Append(
		segment.Segment{
			ID:   1,
			Data: []byte("segment"),
		},
	)

	fmt.Println(
		"graph:",
		graph.Size(),
	)

	fmt.Println(
		"log:",
		store.Size(),
	)

	fmt.Println(
		"gc:",
		gc.Cycles(),
	)

	fmt.Println(
		"epoch:",
		epochValue,
	)

	fmt.Println(
		"segments:",
		segments.Size(),
	)
}
