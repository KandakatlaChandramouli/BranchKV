package main

import (
	"branchkv-core/internal/unsafe_mem"
	"fmt"
)

func main() {

	page := unsafe_mem.NewRawPage(1024)

	page.WriteFloat32(0, 123.456)

	v := page.ReadFloat32(0)

	fmt.Println("Unsafe Value:", v)
}
