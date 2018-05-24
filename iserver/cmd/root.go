package cmd

var cfgFile string
var logFile string
var dbFile string

type ServerExit interface {
	Stop()
}

var serverExit []ServerExit


func exitLoop() {
	exit := make(chan bool)
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	defer signal.Stop(c)
	defer close(exit)

	go func() {
		<-c
		fmt.Printf("IOST server received interrupt, shutting down...")

		for _, s := range serverExit {
			if s != nil {
				s.Stop()
			}
		}

		os.Exit(0)
	}()

	<-exit
}
