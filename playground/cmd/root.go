package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/spf13/viper"
	"Go-IOS-Protocol/verifier"
	"Go-IOS-Protocol/vm"
	"Go-IOS-Protocol/vm/lua"
	"fmt"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "playground",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		db := Database{make(map[string][]byte)}
		mdb := state.NewDatabase(&db)
		pool := state.NewPool(mdb)
		for _, k := range viper.AllKeys() {
			v := viper.GetString(k)
			val, _ := state.ParseValue(v)
			pool.Put(state.Key(k), val)
		}

		v := verifier.NewCacheVerifier(pool)
		var sc0 vm.Contract

		switch language {
		case "lua":
			for i, file := range args {
				code := ReadSourceFile(file)
				parser, err := lua.NewDocCommentParser(code)
				if err != nil {
					panic(err)
				}
				sc, err := parser.Parse()
				if err != nil {
					panic(err)
				}
				if i == 0 {
					sc0 = sc
				}

				v.StartVM(sc)
			}
		default:
			fmt.Println(language, "not supported")
		}
		pool2, gas, err := v.Verify(sc0)
		if err != nil {
			fmt.Println("error:", err.Error())
		}
		pool2.Flush()
		fmt.Println("======Report")
		fmt.Println("gas spend:", gas)
		fmt.Println("state trasition:")
		for k, v := range db.Normal {
			fmt.Printf("  %v: %v\n", k, string(v))
		}

	},

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("./values.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

}