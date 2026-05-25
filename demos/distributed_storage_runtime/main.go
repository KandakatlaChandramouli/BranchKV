package main

import (
	"branchkv-core/internal/distributedwal"
	"branchkv-core/internal/objectcache"
	"branchkv-core/internal/queryplanner"
	"branchkv-core/internal/raftlog"
	"branchkv-core/internal/segmenttree"
	"branchkv-core/internal/txn"
	"branchkv-core/internal/vectorindex"
	"branchkv-core/internal/writebatch"
	"fmt"
)

func main() {

	log := raftlog.NewRaftLog()

	log.Append(
		1,
		[]byte("raft"),
	)

	index := vectorindex.NewIndex()

	index.Insert(
		1,
		[]float32{1, 2, 3},
	)

	plan := queryplanner.NewPlan()

	plan.Add(
		"scan",
	)

	cache := objectcache.NewCache()

	cache.Put(
		"runtime",
		[]byte("healthy"),
	)

	tx := txn.Begin(1)

	tx.Commit()

	batch := writebatch.NewBatch()

	batch.Put(
		"kv",
		[]byte("branch"),
	)

	tree := segmenttree.NewTree()

	tree.Add(1)

	wal := distributedwal.NewDistributedWAL()

	wal.AddReplica(1)

	fmt.Println("raft:", log.Size())
	fmt.Println("vectors:", index.Size())
	fmt.Println("plan:", plan.Steps())
	fmt.Println("batch:", batch.Size())
	fmt.Println("segments:", tree.Size())
	fmt.Println("replicas:", wal.Replicas())
}
