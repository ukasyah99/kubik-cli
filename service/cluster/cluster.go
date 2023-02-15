package cluster

import "fmt"

type CreateParams struct {
	Name     string
	Provider struct {
		ID    string
		Token string
	}
}

func Create() {
	fmt.Println("TODO: validate input")
	fmt.Println("TODO: store initial data (status: pending)")
}
