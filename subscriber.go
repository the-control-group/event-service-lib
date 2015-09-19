package lib

import (
	"github.com/apcera/nats"
	"time"
)

func Subscribe(nc *nats.Conn, subscriberPrefix, queueGroup string, handler nats.MsgHandler) (subscription *nats.Subscription) {
	for subscription == nil {
		select {
		case <-Done:
			return
		default:
			// Listen to all events so that the interesting events can be changed without un/resubscribing
			var queue = subscriberPrefix + "*.>"
			var err error
			if Log != nil {
				Log.Info("Subscribing to events")
			}
			subscription, err = nc.QueueSubscribe(queue, queueGroup, handler)
			if err != nil {
				if Log != nil {
					Log.Error("Unable to subscribe to events. Retrying in 5s.", err)
				}
				time.Sleep(5 * time.Second)
				continue
			}
			Log.Info("Subscribed to events")
			break
		}
	}
	return
}
