package console

func Exit() Cmd {
	e := Cmd{
		Name:  "exit",
		Usage: `Stop daemon and quit`,
	}

	return e
}
