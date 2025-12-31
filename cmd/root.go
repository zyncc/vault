package cmd

import (
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Secure Local Password Manager",
	Long:  `Local & Secure Password Manager`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label:    "Password Vault",
			Items:    []string{"Open Vault", "Find Password", "Insert into Vault", "Generate"},
			HideHelp: false,
		}

		_, selected, err := prompt.Run()
		if err != nil {
			return
		}

		switch selected {
		case "Open Vault":
			openVaultCommand()
		case "Find Password":
			findPasswordCommand()
		case "Insert into Vault":
			insertCommand()
		case "Generate":
			generateCommand()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
