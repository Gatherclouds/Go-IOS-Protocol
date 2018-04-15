package protocol

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/iost-official/Go-IOS-Protocol/iosbase"
)

//go:generate mockgen -destination mocks/mock_component.go -package protocol_mock github.com/iost-official/Go-IOS-Protocol/protocol Component

type Phase int

const (
	StartPhase Phase = iota
	PrePreparePhase
	PreparePhase
	CommitPhase
	PanicPhase
	EndPhase
)

const (
	PrePrepareTimeout = 58 * time.Second
	PrepareTimeout    = 1 * time.Second
	CommitTimeout     = 1 * time.Second
)

type Component interface {
	Init(self iosbase.Member, db Database, router Router) error
	Run()
	Stop()
}

