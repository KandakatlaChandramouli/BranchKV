package specwalk

import (
	"branchkv-core/internal/virtual_mem/pagewalker"
	"branchkv-core/internal/virtual_mem/snapshotspace"
	"branchkv-core/internal/virtual_mem/tlb"
)

type Walker struct {
	space *snapshotspace.Space
	cache *tlb.Cache
}

func NewWalker(
	space *snapshotspace.Space,
	cache *tlb.Cache,
) *Walker {

	return &Walker{
		space: space,
		cache: cache,
	}
}

func (w *Walker) Inner() *pagewalker.Walker {

	return pagewalker.NewWalker(
		w.space.Root,
		w.cache,
	)
}
