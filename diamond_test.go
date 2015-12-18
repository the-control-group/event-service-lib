package lib

import (
	"os"
	"strings"
	"testing"
)

func TestGetDiamonHostname(t *testing.T) {
	var hostname, _ = os.Hostname()
	t.Log(hostname)
	var diamondHostname = GetDiamondHostname()
	t.Log(diamondHostname)
	if strings.Contains(diamondHostname, ".") {
		t.Error("Diamond hostname should just be the first part of the OS hostname")
	}
}
