package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/cmd/credential"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update new resource",
	Long:  "Update new resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Update called")
	},
}

func init() {
	updateCmd.AddCommand(credential.UpdateCredentialCmd)
	rootCmd.AddCommand(updateCmd)
}
