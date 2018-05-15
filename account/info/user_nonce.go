package info

type UserNonce struct {
	UserInfo
}

const (
	UserNonce_Encoding_Head        = "[USER'S NONCE] "
	UserNonce_Encoding_Tail        = "[/USER'S NONCE]"
	UserNonce_Encoding_Head_Length = len(UserNonce_Encoding_Head)
	UserNonce_Encoding_Tail_Length = len(UserNonce_Encoding_Tail)
)

