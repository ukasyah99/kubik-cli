package credential

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/repository/credential"
)

var CreateCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Create new credential",
	Long:  "Create new credential",
	Run: func(cmd *cobra.Command, args []string) {
		var questions = []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: "Enter name:"},
				Validate: survey.Required,
			},
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

		answers := struct {
			Name     string
			Provider string
			Token    string
		}{}

		if err := survey.Ask(questions, &answers); err != nil {
			fmt.Println(err.Error())
			return
		}

		cred := credential.Credential{}
		cred.Name = answers.Name
		cred.Provider = answers.Provider
		cred.Token = answers.Token

		if err := cred.Save(); err != nil {
			fmt.Println(err.Error())
		}
	},
}
