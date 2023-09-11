//go:build windows

package oneinstance

import (
	"net"

	npipe "github.com/natefinch/npipe"
)

func dial() (net.Conn, error) {
	return npipe.Dial(`\\.\pipe\changeme`)
}

func listen() (net.Listener, error) {
	return npipe.Listen(`\\.\pipe\changeme`)
}
