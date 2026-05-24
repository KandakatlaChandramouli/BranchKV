package main

import (
	"branchkv-core/internal/checksum"
	"branchkv-core/internal/hypervisor"
	"branchkv-core/internal/interrupt"
	"branchkv-core/internal/memorymap"
	"branchkv-core/internal/process"
	"branchkv-core/internal/scheduler_rt"
	"branchkv-core/internal/syscall"
	"branchkv-core/internal/threading"
	"fmt"
)

func main() {

	hyper := hypervisor.NewHypervisor()

	vm := hyper.CreateVM()

	sys := syscall.NewTable()

	sys.Register(
		1,
		func() error {
			return nil
		},
	)

	_ = sys.Invoke(1)

	irq := interrupt.NewInterruptController()

	irq.Trigger()

	pm := process.NewProcessManager()

	proc := pm.Spawn()

	tm := threading.NewThreadManager()

	thread := tm.Spawn()

	mmap := memorymap.NewMemoryMap()

	mmap.Map(
		4096,
		8192,
	)

	rt := scheduler_rt.NewRTScheduler()

	rt.Schedule()

	cs := checksum.NewChecksum()

	sum := cs.Compute(
		[]byte("runtime"),
	)

	fmt.Println(
		"vm:",
		vm.ID,
	)

	fmt.Println(
		"irq:",
		irq.Count(),
	)

	fmt.Println(
		"process:",
		proc.PID,
	)

	fmt.Println(
		"thread:",
		thread.TID,
	)

	fmt.Println(
		"mmap:",
		mmap.Size(),
	)

	fmt.Println(
		"rt:",
		rt.Count(),
	)

	fmt.Println(
		"checksum:",
		sum,
	)
}
