// +build windows

package client

import (
	"net"

	"github.com/Microsoft/go-winio"
)

func newDirectPipeConn(c *Client, network, address string) (net.Conn, error) {
	return winio.DialPipe(address, &c.option.ConnectTimeout)
}
