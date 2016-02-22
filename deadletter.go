package lib

import (
	"encoding/json"
	"gopkg.in/Sirupsen/logrus.v0"
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
	Subject        string    `json:"subject"` // The name of the event
	Reason         string    `json:"reason"` // The generic reason for the deadletter (i.e. "database error")
	Message        string    `json:"message"` // A more qualified reason for the deadletter (i.e. "commit failed")
	Process        string    `json:"process"` // The name of the process that generated the deadletter so it can be traced and backfilled
	ProcessVersion string    `json:"process_version"` // The version of the process that generated the deadletter
	Errors         []string  `json:"errors"` // A place to include any raw errors that triggered the deadletter
	Created        time.Time `json:"created"`
}

func SendDeadletter(log *logrus.Entry, pub Publisher, subject, reason, message, process, process_version string, errors []string) {
	dl := Deadletter{subject, reason, message, process, process_version, errors, time.Now()}
	log.WithFields(logrus.Fields{"deadletter": dl}).Warn("Publishing deadletter message")
	var dlJson, err = json.Marshal(dl)
	if err != nil {
		log.WithError(err).WithField("deadletter", dl).Error("Unable to marshal deadletter json")
		return
	}
	err = pub.Publish("deadletter", dlJson)
	if err != nil {
		log.WithError(err).Error("Unable to publish deadletter message")
	}
}
