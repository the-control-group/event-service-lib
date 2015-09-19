package lib

import (
	"github.com/apcera/nats"
	"time"
)

var subscription *nats.Subscription

func Subscribe(nc *nats.Conn, subscriberPrefix, queueGroup string, handler nats.MsgHandler) {
	for {
		select {
		case <-Done:
			subscription.Unsubscribe()
			return
		default:
			for !nc.IsClosed() && (subscription == nil || !subscription.IsValid()) {
				select {
				case <-Done:
					subscription.Unsubscribe()
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
				}
			}
			time.Sleep(5 * time.Second)
		}
	}
}
