// wrap of binary/byte level operates

/*
This package contains structures of IOS block-chain and functions to operate them
*/
package iostdb

type Serializable interface {
	Encode() []byte
	Decode([]byte) error
	Hash() []byte
}

type Binary struct {
	bytes  []byte
	length int
}

