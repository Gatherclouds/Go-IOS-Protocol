package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.iserver.yaml)")
	rootCmd.PersistentFlags().StringVar(&logFile, "log", "", "log file (default is ./iserver.log)")
	rootCmd.PersistentFlags().StringVar(&dbFile, "db", "", "database file (default is ./data.db)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("log", rootCmd.PersistentFlags().Lookup("log"))
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

