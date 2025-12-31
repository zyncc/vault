package cmd

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/zyncc/vault/db"
	"github.com/zyncc/vault/password"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the vault to show all passwords",
	Run: func(cmd *cobra.Command, args []string) {
		openVaultCommand()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func openVaultCommand() {
	enterMasterPrompt := promptui.Prompt{
		Label: "Enter Master Password",
		Mask:  '*',
	}
	masterPass, err := enterMasterPrompt.Run()
	if err != nil {
		return
	}

	q := db.Init()

	rawMasterQuery, err := q.GetMasterPassword(context.Background())
	if err != nil {
		return
	}

	valid, err := password.CompareHash(masterPass, rawMasterQuery.Password, rawMasterQuery.Salt)
	if err != nil {
		fmt.Println("couldnt compare master password")
		return
	}
	if !valid {
		fmt.Println("\n❌ Invalid Master Password!")
		return
	}

	passwords, err := q.GetAllPasswords(context.Background())
	if err != nil {
		return
	}

	fmt.Println(renderTable(passwords))
}
