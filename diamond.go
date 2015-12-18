package lib

import (
	"os"
	"strings"
)

func GetDiamondHostname() (hostname string) {
	hostname, _ = os.Hostname()
	hostname = strings.Split(hostname, ".")[0]
	return
}
