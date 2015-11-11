package lib

import (
	"encoding/json"
	"gopkg.in/Sirupsen/logrus.v0"
	"github.com/the-control-group/nats"
	"time"
)

var DLS_REASON_UNKNOWN = "Unknown reason"
var DLS_REASON_JSON_DECODE = "Unable to decode json"
var DLS_REASON_JSON_ENCODE = "Unable to encode json"
var DLS_REASON_UNRECOGNIZED_TYPE = "Unrecognized message type"
var DLS_REASON_INVALID_EVENT = "Invalid event"
var DLS_REASON_WRITE_FAILED = "Write failed"
var DLS_REASON_MALFORMED_SUBJECT = "Malformed subject"

type Deadletter struct {
	Subject string    `json:"subject"`
	Reason  string    `json:"reason"`
	Message string    `json:"message"`
	Process string    `json:"process"`
	Errors  []string  `json:"errors"`
	Created time.Time `json:"created"`
}

func SendDeadletter(log *logrus.Entry, nc *nats.Conn, subject, reason, message, process string, errors []string) {
	dl := Deadletter{subject, reason, message, process, errors, time.Now()}
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
