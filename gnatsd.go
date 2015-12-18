package lib

import (
	"github.com/the-control-group/nats"
	"gopkg.in/Sirupsen/logrus.v0"
	"net"
)

func NewGnatsConnection(log *logrus.Entry, addr Address) (nc *nats.Conn, err error) {
	var opts = nats.DefaultOptions
	opts.Servers = []string{"nats://" + net.JoinHostPort(addr.Host, addr.Port)}
	opts.ClosedCB = func(nc *nats.Conn) {
		log.Warn("Nats connection closed")
	}
	opts.DisconnectedCB = func(nc *nats.Conn) {
		log.Warn("Nats connection disconnected")
	}
	opts.ReconnectedCB = func(nc *nats.Conn) {
		log.Info("Nats connection reconnected")
	}
	opts.AsyncErrorCB = func(nc *nats.Conn, sub *nats.Subscription, err error) {
		log.WithError(err).Error("Nats Async Error")
	}
	nc, err = opts.Connect()
	return
}
