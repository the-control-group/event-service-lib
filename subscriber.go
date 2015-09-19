package lib

import (
	"github.com/apcera/nats"
)

func Subscribe(log logger, nc *nats.Conn, subscriberPrefix, queueGroup string, handler nats.MsgHandler) (subscription *nats.Subscription, err error) {
	// Listen to all events so that the interesting events can be changed without un/resubscribing
	var queue = subscriberPrefix + "*.>"
	log.Info("Subscribing to events")
	subscription, err = nc.QueueSubscribe(queue, queueGroup, handler)
	if err != nil {
		return
	}
	log.Info("Subscribed to events")
	return
}
