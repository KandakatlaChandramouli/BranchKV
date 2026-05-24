package rpc

import (
	"net"
	"net/rpc"
)

type Request struct {
	Value int
}

type Response struct {
	Result int
}

type RuntimeRPC struct{}

func (r *RuntimeRPC) Double(
	req Request,
	resp *Response,
) error {

	resp.Result = req.Value * 2

	return nil
}

func StartServer(
	addr string,
) error {

	server := rpc.NewServer()

	err := server.Register(
		&RuntimeRPC{},
	)

	if err != nil {
		return err
	}

	listener, err := net.Listen(
		"tcp",
		addr,
	)

	if err != nil {
		return err
	}

	go func() {

		for {

			conn, err := listener.Accept()

			if err != nil {
				continue
			}

			go server.ServeConn(conn)
		}
	}()

	return nil
}

func Call(
	addr string,
	v int,
) (int, error) {

	client, err := rpc.Dial(
		"tcp",
		addr,
	)

	if err != nil {
		return 0, err
	}

	defer client.Close()

	req := Request{
		Value: v,
	}

	var resp Response

	err = client.Call(
		"RuntimeRPC.Double",
		req,
		&resp,
	)

	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}
