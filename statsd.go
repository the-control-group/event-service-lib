package lib

import (
	"github.com/cactus/go-statsd-client/statsd"
	"net"
	"strings"
	"time"
)

const StatsSampleRate = 1

func NewStatsdBuffer(c Emitter, hostname, serviceName string) (statsdbuffer statsd.Statter, err error) {
	if hostname == "" {
		hostname = "EventsServiceUnknownHost"
	}
	hostname = CleanStatsdComponent(hostname)
	if serviceName == "" {
		serviceName = "EventsServiceUnknownService"
	}
	serviceName = CleanStatsdComponent(serviceName)
	prefix := strings.Join([]string{c.Prefix, serviceName, hostname}, ".")
	prefix = strings.TrimLeft(prefix, ".")
	statsdbuffer, err = statsd.NewBufferedClient(net.JoinHostPort(c.Host, c.Port), prefix, time.Duration(c.Interval)*time.Second, 0)
	if err != nil {
		return
	}
	return
}

// Removes characters with special meaning from event components being sent to statsd
func CleanStatsdComponent(name string) string {
	return strings.Replace(strings.Replace(name, ":", "_", -1), ".", "_", -1)
}

func StatsdEventName(parts ...string) string {
	return strings.Join(parts, ".")
}

// Emit received events as `event.$event-name`
func EmitEventReceived(stats statsd.Statter, eventName string) (err error) {
	eventName = CleanStatsdComponent(eventName)
	if stats != nil {
		err = stats.Inc(StatsdEventName("event", eventName), 1, StatsSampleRate)
	}
	return
}

// Emit invalid events as `invalid.$event-name`
func EmitEventInvalid(stats statsd.Statter, eventName string) (err error) {
	eventName = CleanStatsdComponent(eventName)
	if stats != nil {
		err = stats.Inc(StatsdEventName("invalid", eventName), 1, StatsSampleRate)
	}
	return
}

// Emit errors as `error.$event-name`
func EmitError(stats statsd.Statter, eventName string) (err error) {
	eventName = CleanStatsdComponent(eventName)
	if stats != nil {
		err = stats.Inc(StatsdEventName("error", eventName), 1, StatsSampleRate)
	}
	return
}

// Emit failed actions as `$action.failure.$event-name`
func EmitActionFailure(stats statsd.Statter, action, eventName string) (err error) {
	action = CleanStatsdComponent(action)
	eventName = CleanStatsdComponent(eventName)
	if stats != nil {
		err = stats.Inc(StatsdEventName("failure", action, eventName), 1, StatsSampleRate)
	}
	return
}

// Emit successful actions as `$action.success.$event-name`
func EmitActionSuccess(stats statsd.Statter, action, eventName string) (err error) {
	action = CleanStatsdComponent(action)
	eventName = CleanStatsdComponent(eventName)
	if stats != nil {
		err = stats.Inc(StatsdEventName("success", action, eventName), 1, StatsSampleRate)
	}
	return
}

// Emit successful actions as `$action.success.$event-name`
func EmitActionSuccessTiming(stats statsd.Statter, action, eventName string, took time.Duration) (err error) {
	action = CleanStatsdComponent(action)
	eventName = CleanStatsdComponent(eventName)
	if stats != nil {
		err = stats.TimingDuration(StatsdEventName("success", eventName), took, StatsSampleRate)
	}
	return
}

// Emit stop signals as `stop.$signal`
func EmitStopSignal(stats statsd.Statter, signal string) (err error) {
	signal = CleanStatsdComponent(signal)
	if stats != nil {
		err = stats.Inc(StatsdEventName("stop", signal), 1, StatsSampleRate)
	}
	return
}

// Emit reloads as `reload`
func EmitReload(stats statsd.Statter) (err error) {
	if stats != nil {
		err = stats.Inc(StatsdEventName("reload"), 1, StatsSampleRate)
	}
	return
}
