package cmd

type Database struct {
	Normal map[string][]byte
}

func (d *Database) Put(key []byte, value []byte) error {
	d.Normal[string(key)] = value
	return nil
}

func (d *Database) PutHM(key []byte, args ...[]byte) error {
	key1 := string(key)
	key2 := string(args[0])
	d.Normal[key1+"."+key2] = args[1]
	return nil
}
func (d *Database) Get(key []byte) ([]byte, error) {
	return d.Normal[string(key)], nil
}
