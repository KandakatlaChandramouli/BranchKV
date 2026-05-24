package main

import (
	"branchkv-core/internal/mmap"
	"branchkv-core/internal/snapshot"
	"branchkv-core/internal/wal"
	"fmt"
)

type RuntimeState struct {
	Value int
}

func main() {

	w, _ := wal.NewWAL("runtime.wal")

	defer w.Close()

	_ = w.Append(999)

	region, _ := mmap.NewMMapRegion(
		"runtime.mmap",
		4096,
	)

	defer region.Close()

	region.Data[0] = 77

	engine := snapshot.NewSnapshotEngine(
		"runtime_snapshot.gob",
	)

	state := RuntimeState{
		Value: 123,
	}

	_ = engine.Save(state)

	fmt.Println("runtime initialized")
}
