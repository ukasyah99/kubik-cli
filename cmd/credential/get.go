package credential

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/db"
	"github.com/ukasyah99/kubik-cli/db/model"
)

var GetCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Get credential",
	Long:  "Get credential",
	Run: func(cmd *cobra.Command, args []string) {
		data := []model.Credential{}

		db.DB.Find(&data)

		result, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(result))
	},
}
