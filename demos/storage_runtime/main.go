package main

import (
	"branchkv-core/internal/checkpoint"
	"branchkv-core/internal/pagecache"
	"branchkv-core/internal/recoverylog"
	"branchkv-core/internal/sstable"
	"branchkv-core/internal/wal"
	"fmt"
)

func main() {

	w, err := wal.Open(
		"runtime.wal",
	)

	if err != nil {
		panic(err)
	}

	defer w.Close()

	err = w.Append(
		[]byte("branchkv"),
	)

	if err != nil {
		panic(err)
	}

	table := sstable.NewTable()

	table.Add(
		"runtime",
		[]byte("active"),
	)

	cache := pagecache.NewCache()

	cache.Put(
		1,
		[]byte("cached"),
	)

	log := recoverylog.NewLog()

	log.Append(
		1,
		[]byte("recover"),
	)

	snapshot := checkpoint.Snapshot{
		State: map[string]string{
			"runtime": "healthy",
		},
	}

	err = checkpoint.Save(
		"checkpoint.json",
		snapshot,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"sstable:",
		table.Size(),
	)

	fmt.Println(
		"recovery:",
		log.Size(),
	)
}
