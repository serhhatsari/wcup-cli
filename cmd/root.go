package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var flagVerbose bool

var cmdRoot = &cobra.Command{
	Use:   "wcup",
	Short: "wcup CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func Execute() error {
	cmdRoot.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output")

	cmdRoot.MarkPersistentFlagRequired("port")
	//if err := cmdRoot.Execute(); err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}

	return cmdRoot.Execute()

}

func init() {
	// Flags
	//cmdRoot.PersistentFlags().Int16P("port", "p", 8080, "listening port")

	viper.Set("Verbose", true)

	// Set the file name of the configurations file
	viper.SetConfigName("wcup")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	// Set flags
	//viper.BindPFlag("port", cmdRoot.PersistentFlags().Lookup("port"))

	// Set undefined variables
	viper.SetDefault("version", "0.1.0")

	// If config file exists
	//if err := viper.ReadInConfig(); err != nil {
	//	fmt.Println("fatal error config file: default \n", err)
	//	os.Exit(1)
	//}

}
