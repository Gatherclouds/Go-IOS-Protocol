package log

import (
	"os"
	"sync"
)

type Logger struct {
	Tag     string
	logFile *os.File
}

var instance *os.File

var initialized int32
var mu sync.Mutex

const Path = "test.log"
