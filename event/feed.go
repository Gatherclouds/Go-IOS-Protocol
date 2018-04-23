package event

import (
	"sync"
	"reflect"
)

type Feed struct {
	once      sync.Once        // ensures that init only runs once
	sendLock  chan struct{}    // sendLock has a one-element buffer and is empty when held.It protects sendCases.
	removeSub chan interface{} // interrupts Send
	sendCases caseList         // the active set of select cases used by Send

	// The inbox holds newly subscribed channels until they are added to sendCases.
	mu     sync.Mutex
	inbox  caseList
	etype  reflect.Type
	closed bool
}
type caseList []reflect.SelectCase
