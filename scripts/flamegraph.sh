#!/bin/bash

go test -bench=. \
-cpuprofile cpu.prof \
./benchmarks

go tool pprof -http=:8080 cpu.prof
