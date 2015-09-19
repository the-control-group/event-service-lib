package lib

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

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}