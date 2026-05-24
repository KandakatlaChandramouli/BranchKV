package main

import (
	"branchkv-core/internal/branch_predictor"
	"branchkv-core/internal/cache"
	"branchkv-core/internal/metrics"
	"branchkv-core/internal/prefetch"
	"branchkv-core/internal/types"
	"fmt"
)

func main() {

	cacheEngine := cache.NewTokenCache()

	metricsEngine := metrics.NewMetrics()

	prefetcher := prefetch.NewPrefetcher(4)

	predictor := branch_predictor.NewPredictor()

	seq := []types.TokenID{
		10, 20, 30, 40, 50,
	}

	cacheEngine.Store(
		10,
		[]float32{
			1.1,
			2.2,
			3.3,
		},
	)

	window := prefetcher.Predict(seq)

	next := predictor.PredictNext(seq)

	metricsEngine.RecordTrieSearch()

	fmt.Println("prefetch window:", window)

	fmt.Println("predicted token:", next)

	fmt.Println(
		"searches:",
		metricsEngine.TrieSearches.Load(),
	)
}
