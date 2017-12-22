package core

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/cloudfoundry/gosigar"
)

func init() {
	go gchandler()
}

func gchandler() {
	mem := sigar.Mem{}
	for {
		mem.Get()
		if mem.Free < 256*1024*1024 {
			log.Println("[!] Low memory: running garbage collection")
			debug.FreeOSMemory()
		}
		time.Sleep(10 * time.Second)
	}
}
