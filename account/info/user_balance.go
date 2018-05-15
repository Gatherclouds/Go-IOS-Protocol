package info

type UserBalance struct {
	UserInfo
}

const (
	UserBalance_Encoding_Head        = "[USER'S BALANCE]"
	UserBalance_Encoding_Tail        = "[/USER'S BALANCE]"
	UserBalance_Encoding_Head_Length = len(UserBalance_Encoding_Head)
	UserBalance_Encoding_Tail_Length = len(UserBalance_Encoding_Tail)
)
