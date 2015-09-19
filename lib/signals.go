package lib

import (
	"os"
)

var SigChan = make(chan os.Signal, 1)

func init() {
	go signalHandler(SigChan)
}

var SignalHandlers = map[os.Signal][]func(){}

func signalHandler(sigChan chan os.Signal) {
	for {
		sig := <-sigChan
		if fns, ok := SignalHandlers[sig]; ok {
			for _, fn := range fns {
				fn()
			}
		}
	}
}
