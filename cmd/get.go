package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/cmd/credential"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource",
	Long:  "Get resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	getCmd.AddCommand(credential.GetCredentialCmd)
	rootCmd.AddCommand(getCmd)
}
