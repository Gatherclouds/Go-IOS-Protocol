package drp

import (
	"math/big"
)

// globals: block, msg

type Address [32]byte

type Participant struct {
	secret     *big.Int
	commitment [32]byte
	reward     *big.Int
	revealed   bool
	rewarded   bool
}

type Consumer struct {
	address   Address
	bountyPot *big.Int
}

type Campaign struct {
	blockNumber    uint32
	deposit        *big.Int
	commitBalkline uint16
	commitDeadline uint16

	random     *big.Int
	settled    bool
	bountyPot  *big.Int
	numCommits uint32
	numReveals uint32

	consumers    map[Address]*Consumer
	participants map[Address]*Participant
}

// TODO: Persist this to LevelDB
