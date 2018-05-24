package cmd

var cfgFile string
var logFile string
var dbFile string

type ServerExit interface {
	Stop()
}

var serverExit []ServerExit

