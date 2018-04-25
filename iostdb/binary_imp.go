package iostdb

func NewBinary() Binary {
	var b Binary
	b.bytes = []byte{}
	b.length = 0
	return b
}

func BinaryFromBase58(s string) Binary {
	bits := Base58Decode(s)
	return Binary{bits, len(bits)}
}

func (bin *Binary) ToBase58() string {
	return Base58Encode(bin.bytes)
}













