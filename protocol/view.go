package protocol

import (
	"fmt"
	"sort"

	"github.com/iost-official/Go-IOS-Protocol/iosbase"
)

//go:generate mockgen -destination mocks/mock_view.go -package protocol_mock github.com/iost-official/Go-IOS-Protocol/protocol View

/*
Information of PBFT committee members
*/
type View interface {
	Init(chain iosbase.BlockChain)

	GetPrimary() iosbase.Member
	GetBackup() []iosbase.Member
	IsPrimary(ID string) bool
	IsBackup(ID string) bool
	CommitteeSize() int
	ByzantineTolerance() int
}

func ViewFactory(target string) (View, error) {
	switch target {
	case "pob":
		return &PobView{}, nil
	}
	return nil, fmt.Errorf("target view not found")
}


