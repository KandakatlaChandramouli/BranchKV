package tree

import (
	"branchkv-core/internal/types"
	"branchkv-core/internal/virtual_mem"
	"fmt"
	"strings"
)

type RadixTree struct {
	nodes map[string]virtual_mem.VirtualDescriptor
}

func NewRadixTree() *RadixTree {

	return &RadixTree{
		nodes: make(
			map[string]virtual_mem.VirtualDescriptor,
		),
	}
}

func sequenceKey(
	seq []types.TokenID,
) string {

	parts := make(
		[]string,
		0,
		len(seq),
	)

	for _, v := range seq {

		parts = append(
			parts,
			fmt.Sprintf(
				"%d",
				v,
			),
		)
	}

	return strings.Join(
		parts,
		":",
	)
}

func (r *RadixTree) Insert(
	seq []types.TokenID,
	desc virtual_mem.VirtualDescriptor,
) {

	key := sequenceKey(seq)

	r.nodes[key] = desc
}

func (r *RadixTree) Search(
	seq []types.TokenID,
) (
	virtual_mem.VirtualDescriptor,
	bool,
) {

	key := sequenceKey(seq)

	v, ok := r.nodes[key]

	return v,
		ok
}
