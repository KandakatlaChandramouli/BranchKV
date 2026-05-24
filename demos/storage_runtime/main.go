package main

import (
	"branchkv-core/internal/compaction"
	"branchkv-core/internal/filesystem"
	"branchkv-core/internal/io_uring"
	"branchkv-core/internal/objectstore"
	"branchkv-core/internal/raft"
	"branchkv-core/internal/sharding"
	"branchkv-core/internal/sstable"
	"fmt"
)

func main() {

	io := io_uring.NewSubmissionQueue()

	io.Submit()

	fs := filesystem.NewFileSystem()

	fs.Create(
		"runtime",
		[]byte("branchkv"),
	)

	store := objectstore.NewObjectStore()

	store.Put(
		"branch",
		[]byte("state"),
	)

	shard := sharding.NewSharder()

	node := shard.Shard(
		"branch",
		4,
	)

	raftNode := raft.NewRaftNode()

	term := raftNode.NextTerm()

	table := sstable.NewSSTable()

	table.Append(
		"runtime",
		[]byte("kv"),
	)

	compactor := compaction.NewCompactor()

	segment := compactor.Compact(
		[]compaction.Segment{
			{ID: 1},
		},
	)

	fmt.Println(
		"io:",
		io.Count(),
	)

	fmt.Println(
		"shard:",
		node,
	)

	fmt.Println(
		"raft:",
		term,
	)

	fmt.Println(
		"sstable:",
		table.Size(),
	)

	fmt.Println(
		"segment:",
		segment.ID,
	)
}
