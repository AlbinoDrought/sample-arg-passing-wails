//go:build !windows

package oneinstance

import (
	"net"
)

func dial() (net.Conn, error) {
	return net.Dial("tcp", "127.0.0.1:12345")
}

func listen() (net.Listener, error) {
	return net.Listen("tcp", "127.0.0.1:12345")
}
