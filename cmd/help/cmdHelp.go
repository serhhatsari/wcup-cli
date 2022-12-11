package help

import (
	"fmt"

	"github.com/spf13/cobra"
)


func Help(cmd *cobra.Command, args []string) {
	fmt.Println("Help")
}

var CmdHelp = &cobra.Command{
	Use:   "--help",
	Short: "-h",
	Run: Help,
}