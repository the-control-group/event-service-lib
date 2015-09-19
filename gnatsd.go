package lib

import (
	"github.com/apcera/nats"
	"net"
	"time"
)

/*
	Should retry indefinitely to establish connection and should reconnect indefinitely if connection is lost
*/
func newGnatsConnection(addr *Address) *nats.Conn {
	var err error

	var opts = nats.DefaultOptions

	var nc *nats.Conn

	var attempts = 0
	for nc == nil {
		attempts = attempts + 1
		opts.Servers = []string{"nats://" + net.JoinHostPort(addr.Host, addr.Port)}
		nc, err = opts.Connect()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
	}

	return nc
}
