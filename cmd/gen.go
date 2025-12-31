package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Can be used to generate Secure Passwords and UUIDs",
	Run: func(cmd *cobra.Command, args []string) {
		generateCommand()
	},
}

func generateCommand() {
	prompt := promptui.Select{
		Label: "Generate",
		Items: []string{"Password", "UUID"},
	}

	_, selected, err := prompt.Run()
	if err != nil {
		return
	}

	switch selected {
	case "Password":
		password, err := generatePassword(length)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("\n", password)

	case "UUID":
		fmt.Println("\n", uuid.NewString())
	}
}

func init() {
	rootCmd.AddCommand(genCmd)
}
