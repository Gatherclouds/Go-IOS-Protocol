package drp

import (
	"fmt"
	"encoding/hex"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

func (s Scalar) String() string {
	bi := s.toInt()
	return fmt.Sprintf("Scalar %s", bi.String())
}

func (p Point) String() string {
	return fmt.Sprintf("Point {x = %d,y = %d}", p.x, p.y)
}

func (dh DhSecret) String() string {
	return hex.EncodeToString(dh.data)
}

func getCurve() elliptic.Curve {
	return elliptic.P256()
}

func getCurveParams() *elliptic.CurveParams {
	return getCurve().Params()
}

func keypairGen() KeyPair {
	priv, x, y, _ := elliptic.GenerateKey(getCurve(), rand.Reader)
	return KeyPair{Scalar{priv}, Point{x, y}}
}

func (s Scalar) toInt() *big.Int {
	return new(big.Int).SetBytes(s.data)
}

func (s *Scalar) fromInt(bi *big.Int) *Scalar {
	order := getCurveParams().N
	bitSize := getCurveParams().BitSize
	if bi.Cmp(big.NewInt(0)) == -1 {
		bi.Add(bi, order)
	}
	b := bi.Bytes()
	blen := cap(b)
	nbBytes := bitSize / 8
	switch {
	case blen == nbBytes:
		s.data = b
	case blen < nbBytes:
		s.data = make([]byte, nbBytes)
		copy(s.data[nbBytes-blen:], b)
	}
	return s
}

func (s *Scalar) fromSmallInt(i int) *Scalar {
	bi := big.NewInt(int64(i))
	return s.fromInt(bi)
}








