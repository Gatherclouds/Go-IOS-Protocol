package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/spf13/viper"
	"Go-IOS-Protocol/verifier"
	"Go-IOS-Protocol/vm"
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

	},

}
