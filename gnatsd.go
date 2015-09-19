package lib

import (
	"github.com/apcera/nats"
	"net"
)

/*
	Should retry indefinitely to establish connection and should reconnect indefinitely if connection is lost
*/
func NewGnatsConnection(_ logger, addr *Address) (nc *nats.Conn, err error) {
	var opts = nats.DefaultOptions
	opts.Servers = []string{"nats://" + net.JoinHostPort(addr.Host, addr.Port)}
	nc, err = opts.Connect()
	return
}
