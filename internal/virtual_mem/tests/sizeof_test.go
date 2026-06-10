package tests

import (
    "fmt"
    "testing"
    "unsafe"

    "branchkv-core/internal/virtual_mem/dagnodearena"
    "branchkv-core/internal/virtual_mem/fixeddag"
)

func TestSizes(t *testing.T) {
    fmt.Println("ArenaNode =", unsafe.Sizeof(dagnodearena.Node{}))
    fmt.Println("FixedNode =", unsafe.Sizeof(fixeddag.Node{}))
}
