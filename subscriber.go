package lib

import (
	"github.com/apcera/nats"
)

func Subscribe(_ logger, nc *nats.Conn, subject, queueGroup string, handler nats.MsgHandler) (subscription *nats.Subscription, err error) {
	// Listen to all events so that the interesting events can be changed without un/resubscribing
	var subject = subject
	subscription, err = nc.QueueSubscribe(subject, queueGroup, handler)
	if err != nil {
		return
	}
	return
}
