package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VERSION = "unversioned"
var COMMITID = "unknown commit"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ara",
	Long:  `All software has versions. This is ara's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ara", VERSION, COMMITID)
	},
}
