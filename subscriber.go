package lib

import (
	"github.com/Sirupsen/logrus"
	"github.com/the-control-group/nats"
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
	var q int
	wg.Add(1)
	defer wg.Done()
	for {
		if !subscription.IsActive() {
			q, err = subscription.QueuedMsgs()
			if err != nil {
				log.Warn(err)
			}
			if q <= 0 {
				break
			}
		}
		msg, err = subscription.NextMsg(1 * time.Second)
		if err != nil {
			switch err {
			case nats.ErrTimeout:
				// Ignore
			case nats.ErrNoMessages:
				break
			default:
				log.WithError(err).Warn("Unable to read next message")
			}
			time.Sleep(1 * time.Second)
			continue
		}
		go handler(msg)
	}
	log.Debug("Done handling messages")
}
