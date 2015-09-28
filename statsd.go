package lib

import (
	"github.com/quipo/statsd"
	"net"
	"os"
	"strings"
	"time"
)

func NewStatsdBuffer(c Emitter, serviceName string) (statsdbuffer *statsd.StatsdBuffer, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "EventsServiceUnknownHost"
	}
	hostname = strings.Replace(hostname, ".", "_", -1)
	statsdClient := statsd.NewStatsdClient(net.JoinHostPort(c.Host, c.Port), strings.Join([]string{c.Prefix, serviceName, hostname}, ".")+".")
	err = statsdClient.CreateSocket()
	if err != nil {
		return
	}
	statsdbuffer = statsd.NewStatsdBuffer(time.Duration(c.Interval)*time.Millisecond, statsdClient)
	return
}
