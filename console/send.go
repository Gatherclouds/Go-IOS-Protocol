package console


func Send() Cmd {
	return Cmd{
		Name:  "send",
		Usage: `send FROM TO AMOUNT - Send AMOUNT of coins from FROM address to TO`,
		Exec: func(host *Console, args []string) string {
			if len(args) != 3 {
				return "Invalid arguments!\n"
			}

			if err != nil {
				return "Invalid arguments!\n"
			}
			return to_print
		},
	}
}
