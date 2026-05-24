package main

import (
	"branchkv-core/internal/wal"
	"fmt"
)

func main() {

	w, err := wal.NewWAL(
		"runtime.wal",
	)

	if err != nil {
		panic(err)
	}

	defer w.Close()

	err = w.Append(
		[]byte("runtime"),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"runtime wal healthy",
	)
}
