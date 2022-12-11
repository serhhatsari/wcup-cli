package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"wcup/cmd/status"
	"wcup/cmd/groups"
	"wcup/cmd/help"
)

var cmdRoot = &cobra.Command{
	Use:   "wcup",
	Short: "wcup CLI",
	Run: help.Help,
}

func Execute() error {

	cmdRoot.MarkPersistentFlagRequired("port")

	groups.CmdGroups.PersistentFlags().StringVarP(&groups.FlagGroup, "group", "g", "", "Group Name")
	status.CmdStatus.PersistentFlags().StringVarP(&status.FlagCountry, "country", "c", "", "Country Name")

	cmdRoot.AddCommand(help.CmdHelp)
	cmdRoot.AddCommand(status.CmdStatus)
	cmdRoot.AddCommand(groups.CmdGroups)

	return cmdRoot.Execute()

}

func init() {

	viper.Set("Verbose", true)

	// Set the file name of the configurations file
	viper.SetConfigName("wcup")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	
	viper.SetDefault("version", "0.1.0")

}
