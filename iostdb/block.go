package iostdb

func (d *Block) Encode() []byte {
	bin, err := d.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return bin
}




