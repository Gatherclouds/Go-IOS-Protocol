package log

import (
	"os"
	"sync"
	"sync/atomic"
	"runtime"
)

type Logger struct {
	Tag     string
	logFile *os.File
}

var instance *os.File

var initialized int32
var mu sync.Mutex

const Path = "test.log" // TODO : 在命令行内修改

func NewLogger(tag string) (*Logger, error) {

	if atomic.LoadInt32(&initialized) == 1 {
		return &Logger{
			logFile: instance,
		}, nil
	}

	mu.Lock()
	defer mu.Unlock()


	runtime.SetFinalizer(instance, func(obj *os.File) {
		obj.Close()
	})

	atomic.StoreInt32(&initialized, 1)

	return &Logger{
		logFile: instance,
	}, nil

}