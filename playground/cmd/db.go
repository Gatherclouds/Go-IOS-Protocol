package cmd

type Database struct {
	Normal map[string][]byte
}

func (d *Database) Put(key []byte, value []byte) error {
	d.Normal[string(key)] = value
	return nil
}
