package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/zyncc/vault/db"
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
			enterMasterPrompt := promptui.Prompt{
				Label: "Enter Master Password",
				Mask:  '*',
			}
			masterPass, err := enterMasterPrompt.Run()
			if err != nil {
				return
			}

			if masterPass != "Master" {
				fmt.Println("Master Password does not match")
			}

			q := db.Init()
			passwords, err := q.GetAllPasswords(context.Background())
			if err != nil {
				return
			}

			for i := range len(passwords) {
				fmt.Printf("\n%v        %v", passwords[i].Domain, passwords[i].Password)
			}

		case "Find Password":
			p1 := promptui.Prompt{
				Label: "Enter Master Password",
				Mask:  '*',
			}
			masterPass, err := p1.Run()
			if err != nil {
				return
			}
			if masterPass != "Master" {
				fmt.Println("\n❌ Master Password does not Match!")
				return
			}

			p2 := promptui.Prompt{
				Label: "Enter the Domain",
			}
			domain, err := p2.Run()
			if err != nil {
				return
			}

			q := db.Init()
			res, err := q.FindPasswordUsingDomain(context.Background(), domain)
			if err != nil {
				return
			}

			fmt.Println("\nPassword: ", res.Password)

		case "Insert into Vault":
			p1 := promptui.Prompt{
				Label: "Enter Domain / URL",
			}
			domain, err := p1.Run()
			if err != nil {
				return
			}
			parseDomain(&domain)

			p3 := promptui.Prompt{
				Label: "Email",
			}
			email, err := p3.Run()
			if err != nil {
				fmt.Println(email)
				return
			}

			p2 := promptui.Prompt{
				Label: "Enter Password",
				Mask:  '*',
			}
			password, err := p2.Run()
			if err != nil {
				return
			}

			if err := insertPasswordToDatabase(&domain, &password); err != nil {
				panic(err.Error())
			}

			fmt.Printf("\n✅ Succesfully added password for %v to the vault\n", domain)
		case "Generate":
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
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
