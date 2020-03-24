package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION is ara's current version
var VERSION = "unversioned"

// COMMITID is populated upon build
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
