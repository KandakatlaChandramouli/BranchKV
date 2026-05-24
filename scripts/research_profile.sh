#!/bin/bash

echo "=== CPU PROFILE ==="

go test \
-bench=. \
-cpuprofile cpu.prof \
./benchmarks

echo "=== MEMORY PROFILE ==="

go test \
-bench=. \
-memprofile mem.prof \
./benchmarks

echo "=== TOP CPU ==="

go tool pprof -top cpu.prof

echo "=== TOP MEMORY ==="

go tool pprof -top mem.prof
