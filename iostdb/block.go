package iostdb

func (d *Block) Encode() []byte {
	bin, err := d.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return bin
}

func (d *Block) Decode(bin []byte) error {
	_, err := d.Unmarshal(bin)
	return err
}



