package zkteco

import (
	"log"
)

var (
	// DEBUG is debug mode. It'll output debug messages if it's true.
	DEBUG = true
)

func debugPrintf(format string, values ...interface{}) {
	if DEBUG {
		log.Printf("[jsondb-debug] "+format, values...)
	}
}
