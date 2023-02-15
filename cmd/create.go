package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/cmd/credential"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new resource",
	Long:  "Create new resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	createCmd.AddCommand(credential.CreateCredentialCmd)
	rootCmd.AddCommand(createCmd)
}
