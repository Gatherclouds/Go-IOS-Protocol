package drp

import (
	"math/big"
)

type Scalar struct {
	data []byte
}

type Point struct {
	x *big.Int
	y *big.Int
}

type DhSecret struct {
	data []byte
}


