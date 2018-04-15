package iosbase

func (d *TextInput) Encode() []byte {
	bin, err := d.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return bin
}

func (d *TextInput) Decode(bin []byte) error {
	_, err := d.Unmarshal(bin)
	return err
}
func (d *TextInput) Hash() []byte {
	return Sha256(d.Encode())
}

func (d *Text) Encode() []byte {
	bin, err := d.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return bin
}

func (d *Text) Decode(bin []byte) error {
	_, err := d.Unmarshal(bin)
	return err
}
func (d *Text) Hash() []byte {
	return Sha256(d.Encode())
}
