#!/bin/bash

go test -bench=. \
-benchmem \
-cpuprofile cpu.prof \
-memprofile mem.prof \
./benchmarks

go tool pprof -top cpu.prof
