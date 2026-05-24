package main

import (
	"branchkv-core/internal/branchcache"
	"branchkv-core/internal/executor"
	"branchkv-core/internal/speculative"
	"branchkv-core/internal/tokenizer"
	"branchkv-core/internal/types"
	"fmt"
)

func main() {

	tokenizerEngine := tokenizer.NewTokenizer()

	specEngine := speculative.NewSpeculativeEngine()

	execEngine := executor.NewExecutor()

	cache := branchcache.NewBranchCache()

	tokens := tokenizerEngine.Tokenize(
		"branch kv speculative runtime",
	)

	seq := make(
		[]types.TokenID,
		0,
	)

	for i := range tokens {
		seq = append(
			seq,
			types.TokenID(i+1),
		)
	}

	branches := specEngine.ForkSequence(seq)

	for i, branch := range branches {

		cache.Store(
			uint64(i),
			branch,
		)

		out := execEngine.Execute(
			branch,
		)

		fmt.Println(
			"branch:",
			i,
			"predicted:",
			out,
		)
	}
}
