package lib

import (
	"net"
	"strings"
	"time"
)

type Emitter struct {
	Address
	Interval int    `json:"interval"`
	Prefix   string `json:"prefix"`
}

type Listener struct {
	Address
	Group  string `json:"group"`
	Prefix string `json:"prefix"`
}

type Address struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (addr Address) String() string {
	return net.JoinHostPort(addr.Host, addr.Port)
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(text []byte) (err error) {
	var str = strings.Trim(string(text), `"`)
	d.Duration, err = time.ParseDuration(str)
	return
}
