package iostdb

func (sp *StatePoolImpl) Init() error {
	var err error
	sp.cli, err = redis.Dial(Conn, DBAddr)
	if err != nil {
		return err
	}
	return nil
}

func (sp *StatePoolImpl) Close() error {
	if sp.cli != nil {
		sp.cli.Close()
	}
	return nil
}






