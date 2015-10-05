package lib

import (
	"os"
	"sync"
)

var SigChan = make(chan os.Signal, 1)

func init() {
	go signalHandler(SigChan)
}

var SignalHandlers = map[os.Signal][]func(){}

var SignalLock sync.Mutex

func signalHandler(sigChan chan os.Signal) {
	for {
		sig := <-sigChan
		SignalLock.Lock()
		if fns, ok := SignalHandlers[sig]; ok {
			for _, fn := range fns {
				fn()
			}
		}
		SignalLock.Unlock()
	}
}
