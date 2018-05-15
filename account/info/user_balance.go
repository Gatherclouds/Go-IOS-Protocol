package info

import (
	"strconv"
	"encoding/base32"
)

type UserBalance struct {
	UserInfo
}

const (
	UserBalance_Encoding_Head        = "[USER'S BALANCE]"
	UserBalance_Encoding_Tail        = "[/USER'S BALANCE]"
	UserBalance_Encoding_Head_Length = len(UserBalance_Encoding_Head)
	UserBalance_Encoding_Tail_Length = len(UserBalance_Encoding_Tail)
)

func (balance *UserBalance) Encode(val uint64) string {
	raw := ([]byte)(UserBalance_Encoding_Head + strconv.FormatUint(val, 10) + UserBalance_Encoding_Tail)
	return base32.HexEncoding.EncodeToString(raw) //or other encoding forms
}

