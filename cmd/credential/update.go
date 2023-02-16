package credential

import (
	"github.com/spf13/cobra"
	"github.com/ukasyah99/kubik-cli/db"
	"github.com/ukasyah99/kubik-cli/db/model"
)

var UpdateCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Update existing credential",
	Long:  "Update existing credential",
	Run: func(cmd *cobra.Command, args []string) {
		var credentials []model.Credential
		db.DB.Find(&credentials)

		// selectOptions := []string{}
		// for _, c := range credentials {
		// 	selectOptions = append(selectOptions, fmt.Sprintf("%d - %s", c.ID, c.Title))
		// }

		// cred := ""
		// prompt := &survey.Select{
		// 	Message: "Choose a credential:",
		// 	Options: selectOptions,
		// }
		// survey.AskOne(prompt, &cred)

		// fmt.Printf("%s chosen!\n", cred)

		// var updateQuestions = []*survey.Question{
		// 	{
		// 		Name: "provider",
		// 		Prompt: &survey.Select{
		// 			Message: "Choose provider:",
		// 			Options: []string{"idcloudhost", "linode"},
		// 			Default: "idcloudhost",
		// 		},
		// 	},
		// 	{
		// 		Name:     "token",
		// 		Prompt:   &survey.Password{Message: "Enter token:"},
		// 		Validate: survey.Required,
		// 	},
		// }

		// data := model.Credential{}

		// err := survey.Ask(updateQuestions, &data)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }

		// db.DB.Create(&data)
	},
}
