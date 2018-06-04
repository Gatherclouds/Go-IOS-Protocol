package message

func (d *RequestHeight) Encode() []byte {
	b, err := d.Marshal(nil)
	if err != nil {
		panic(err)
	}

	return b
}

func (d *RequestHeight) Decode(bin []byte) error {
	_, err := d.Unmarshal(bin)
	if err != nil {
		return err
	}

	return nil
}

