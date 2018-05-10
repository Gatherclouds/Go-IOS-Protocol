package drp

import (
	"fmt"
	"encoding/hex"
	"crypto/elliptic"
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











