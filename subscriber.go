package lib

import (
	logrus "gopkg.in/Sirupsen/logrus.v0"
	"gopkg.in/nats-io/nats.v1"
	"sync"
	"time"
)

func Subscribe(_ *logrus.Entry, nc *nats.Conn, subject, queueGroup string) (subscription *nats.Subscription, err error) {
	subscription, err = nc.QueueSubscribeSync(subject, queueGroup)
	if err != nil {
		return
	}
	return
}

func HandleMessages(log *logrus.Entry, wg *sync.WaitGroup, subscription *nats.Subscription, handler nats.MsgHandler) {
	var msg *nats.Msg
	var err error
	wg.Add(1)
	defer wg.Done()
	for {
		msg, err = subscription.NextMsg(1 * time.Second)
		if err != nil {
			switch err {
			case nats.ErrTimeout:
				// Ignore
			default:
				log.WithError(err).Warn("Unable to read next message")
				return
			}
			time.Sleep(1 * time.Second)
			continue
		}
		go handler(msg)
	}
	log.Debug("Done handling messages")
}
