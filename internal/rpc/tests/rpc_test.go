package tests

import (
	"branchkv-core/internal/rpc"
	"testing"
	"time"
)

func TestRPC(
	t *testing.T,
) {

	go rpc.StartServer(
		":9999",
	)

	time.Sleep(
		time.Millisecond * 200,
	)

	out, err := rpc.Call(
		":9999",
		4,
	)

	if err != nil {
		t.Fatal(err)
	}

	if out != 8 {
		t.Fatal("rpc failed")
	}
}
