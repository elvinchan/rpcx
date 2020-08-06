// +build !windows

package client

import (
	"errors"
	"net"
)

func newDirectPipeConn(c *Client, network, address string) (net.Conn, error) {
	return nil, errors.New("named pipe unsupported")
}
