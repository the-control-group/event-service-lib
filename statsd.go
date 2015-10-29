package lib

import (
	"github.com/cactus/go-statsd-client/statsd"
	"net"
	"strings"
	"time"
)

func NewStatsdBuffer(c Emitter, hostname, serviceName string) (statsdbuffer statsd.Statter, err error) {
	if hostname == "" {
		hostname = "EventsServiceUnknownHost"
	}
	hostname = strings.Replace(hostname, ".", "_", -1)
	if serviceName == "" {
		serviceName = "EventsServiceUnknownService"
	}
	serviceName = strings.Replace(serviceName, ".", "_", -1)
	statsdbuffer, err = statsd.NewBufferedClient(net.JoinHostPort(c.Host, c.Port), strings.Join([]string{c.Prefix, serviceName, hostname}, ".")+".", time.Duration(c.Interval)*time.Second, 0)
	if err != nil {
		return
	}
	return
}

// Removes characters with special meaning from event components being sent to statsd
func CleanStatsdComponent(name string) string {
	return strings.Replace(strings.Replace(name, ":", "_", -1), ".", "_", -1)
}
