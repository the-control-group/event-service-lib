package lib

import (
	"time"
)

const TRANSFORM_FAILED = "Transform failed"

type Transformer struct {
	Name      string
	Prefix    string
	Timeout   time.Duration
	Mandatory bool // Whether or not the message should be deadlettered if a failure occurs
}
