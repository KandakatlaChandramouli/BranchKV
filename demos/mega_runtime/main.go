package main

import (
	"branchkv-core/internal/checksum"
	"branchkv-core/internal/columnstore"
	"branchkv-core/internal/consensus"
	"branchkv-core/internal/deadlockdetector"
	"branchkv-core/internal/graphindex"
	"branchkv-core/internal/hotstorage"
	"branchkv-core/internal/leaderlease"
	"branchkv-core/internal/objectstore"
	"branchkv-core/internal/partitioning"
	"branchkv-core/internal/querycache"
	"branchkv-core/internal/replication"
	"branchkv-core/internal/sharding"
	"fmt"
)

func main() {

	column := columnstore.NewRuntime()
	column.Add()

	cache := querycache.NewRuntime()
	cache.Add()

	store := objectstore.NewRuntime()
	store.Add()

	part := partitioning.NewRuntime()
	part.Add()

	shard := sharding.NewRuntime()
	shard.Add()

	replica := replication.NewRuntime()
	replica.Add()

	graph := graphindex.NewRuntime()
	graph.Add()

	locks := deadlockdetector.NewRuntime()
	locks.Add()

	hot := hotstorage.NewRuntime()
	hot.Add()

	hash := checksum.NewRuntime()
	hash.Add()

	cons := consensus.NewRuntime()
	cons.Add()

	lease := leaderlease.NewRuntime()
	lease.Add()

	fmt.Println("column:", column.Size())
	fmt.Println("cache:", cache.Size())
	fmt.Println("store:", store.Size())
	fmt.Println("partition:", part.Size())
	fmt.Println("shard:", shard.Size())
	fmt.Println("replica:", replica.Size())
	fmt.Println("graph:", graph.Size())
	fmt.Println("locks:", locks.Size())
	fmt.Println("hot:", hot.Size())
	fmt.Println("checksum:", hash.Size())
	fmt.Println("consensus:", cons.Size())
	fmt.Println("lease:", lease.Size())
}
