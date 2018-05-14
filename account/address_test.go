package account

import (
	"testing"
	"fmt"
)

func TestGeneration(t *testing.T) {
	tmp_passphrase := "IOST_TEST_NET"
	addr, err := GenerateAddress(tmp_passphrase)
	if err == nil {
		fmt.Println(addr.ToString())
	} else {
		t.Errorf("Bad passphrase!!!")
	}
}
