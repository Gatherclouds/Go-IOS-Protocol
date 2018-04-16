package drp

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

type KeyPair struct {
	private Scalar
	public  Point
}

