package credential

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"

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

		for i := 0; i < len(data); i++ {
			// hide token value by changing all chars to asterisk
			data[i].Token = strings.Repeat("*", utf8.RuneCountInString(data[i].Token))
		}

		result, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(result))
	},
}
