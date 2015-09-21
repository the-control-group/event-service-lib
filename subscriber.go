package lib

import (
	"github.com/apcera/nats"
)

func Subscribe(_ logger, nc *nats.Conn, subject, queueGroup string) (subscription *nats.Subscription, err error) {
	// Listen to all events so that the interesting events can be changed without un/resubscribing
	subscription, err = nc.QueueSubscribeSync(subject, queueGroup)
	if err != nil {
		return
	}
	return
}
