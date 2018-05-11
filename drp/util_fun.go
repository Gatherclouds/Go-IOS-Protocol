package drp

import (
	"fmt"
	"encoding/hex"
)

func (s Scalar) String() string {
	bi := s.toInt()
	return fmt.Sprintf("Scalar %s", bi.String())
}

func (dh DhSecret) String() string {
	return hex.EncodeToString(dh.data)
}












