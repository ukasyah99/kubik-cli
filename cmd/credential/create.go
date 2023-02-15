package credential

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/db"
	"github.com/ukasyah99/kubik-cli/db/model"
)

var qs = []*survey.Question{
	{
		Name: "provider",
		Prompt: &survey.Select{
			Message: "Choose provider:",
			Options: []string{"idcloudhost", "linode"},
			Default: "idcloudhost",
		},
	},
	{
		Name:     "token",
		Prompt:   &survey.Password{Message: "Enter token:"},
		Validate: survey.Required,
	},
}

var CreateCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Create new credential",
	Long:  "Create new credential",
	Run: func(cmd *cobra.Command, args []string) {
		data := model.Credential{}

		err := survey.Ask(qs, &data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		db.DB.Create(&data)
	},
}
