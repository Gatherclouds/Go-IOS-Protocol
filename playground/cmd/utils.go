package cmd

func ReadSourceFile(file string) (code string) {
	buf, err := ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

