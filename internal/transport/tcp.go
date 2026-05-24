package transport

import (
	"net"
)

type TCPTransport struct {
	Listener net.Listener
}

func NewTCPTransport(
	addr string,
) (*TCPTransport, error) {

	ln, err := net.Listen(
		"tcp",
		addr,
	)

	if err != nil {
		return nil, err
	}

	return &TCPTransport{
		Listener: ln,
	}, nil
}

func (t *TCPTransport) Close() error {
	return t.Listener.Close()
}
