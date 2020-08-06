// +build windows

package server

import (
	"net"

	"github.com/Microsoft/go-winio"
)

func init() {
	makeListeners["pipe"] = pipeMakeListener
}

func pipeMakeListener(s *Server, address string) (ln net.Listener, err error) {
	var c winio.PipeConfig
	if s.tlsConfig != nil && len(s.tlsConfig.NextProtos) > 0 {
		c.SecurityDescriptor = s.tlsConfig.NextProtos[0]
	}
	return winio.ListenPipe(address, &c)
}
