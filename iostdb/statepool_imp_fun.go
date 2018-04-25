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

func (sp *StatePoolImpl) Add(state State) error {
	_, err := sp.cli.Do("HMSET", state.Hash(),
		"value", state.Value,
		"script", state.Script,
		"tx_hash", state.BirthTxHash)
	if err != nil {
		return err
	}
	return nil
}

func (sp *StatePoolImpl) Find(stateHash []byte) (State, error) {
	var s State
	reply, err := redis.Values(sp.cli.Do("HMGET", stateHash, "value", "script", "tx_hash"))
	if err != nil {
		return s, err
	}
	_, err = redis.Scan(reply, &s.Value, &s.Script, s.BirthTxHash)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (sp *StatePoolImpl) Del(stateHash []byte) error {
	_, err := sp.cli.Do("DEL", stateHash)
	if err != nil {
		return err
	}
	return nil
}



