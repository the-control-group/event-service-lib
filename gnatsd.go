package lib

import (
	"github.com/the-control-group/nats"
	"net"
)

func NewGnatsConnection(_ logger, addr *Address) (nc *nats.Conn, err error) {
	var opts = nats.DefaultOptions
	opts.Servers = []string{"nats://" + net.JoinHostPort(addr.Host, addr.Port)}
	nc, err = opts.Connect()
	return
}
