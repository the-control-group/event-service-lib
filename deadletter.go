package lib

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/the-control-group/nats"
)

var DLS_REASON_UNKNOWN = "Unknown reason"
var DLS_REASON_JSON_DECODE = "Unable to decode json"
var DLS_REASON_UNRECOGNIZED_TYPE = "Unrecognized message type"

type Deadletter struct {
	Subject string `json:"subject"`
	Reason  string `json:"reason"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func SendDeadletter(log *logrus.Entry, nc *nats.Conn, subject, reason, errStr, message string) {
	dl := Deadletter{subject, reason, errStr, message}
	log.WithFields(logrus.Fields{"deadletter": dl}).Warn("Publishing deadletter message")
	var dlJson, err = json.Marshal(dl)
	if err != nil {
		log.WithError(err).WithField("deadletter", dl).Error("Unable to marshal deadletter json")
		return
	}
	err = nc.Publish("deadletter", dlJson)
	if err != nil {
		log.WithError(err).Error("Unable to publish deadletter message")
	}
}