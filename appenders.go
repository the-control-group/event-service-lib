package lib

const APPEND_FAILED = "Append failed"

type Appender struct {
	Name      string
	Timeout   time.Duration
	Request   string
	Mandatory bool // Whether or not the message should be deadlettered if a failure occurs
}
