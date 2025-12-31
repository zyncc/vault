package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/zyncc/vault/db"
)

var insertCmd = cobra.Command{
	Use:   "insert",
	Short: "Insert a new password into the vault",
	Run: func(cmd *cobra.Command, args []string) {
		insertCommand()
	},
}

func init() {
	rootCmd.AddCommand(&insertCmd)
}

func insertCommand() {
	p1 := promptui.Prompt{
		Label: "Enter Domain / URL",
	}
	domain, err := p1.Run()
	if err != nil {
		return
	}

	if err := validateInput(&domain); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Parse if URL
	parseDomain(&domain)

	p3 := promptui.Prompt{
		Label: "Email",
	}
	email, err := p3.Run()
	if err != nil {
		return
	}
	if err := validateInput(&email); err != nil {
		fmt.Println(err.Error())
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

	if err := validateInput(&password); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := insertPasswordToDatabase(&domain, &email, &password); err != nil {
		if strings.Contains(err.Error(), "2067") {
			fmt.Println("❌ An Existing Password exists for", domain)
			return
		}
		fmt.Println("❌ Failed to add password to vault")
		return
	}

	fmt.Printf("\n✅ Succesfully added password for %v to the vault\n", domain)
}

func insertPasswordToDatabase(domain *string, email *string, password *string) error {
	q := db.Init()
	return q.InsertIntoPasswordStore(context.Background(), db.InsertIntoPasswordStoreParams{
		ID:       uuid.NewString(),
		Domain:   *domain,
		Email:    *email,
		Password: *password,
	})
}

func parseDomain(domain *string) {
	rawUrl, err := url.Parse(*domain)
	if err != nil {
		return
	}
	hostName := rawUrl.Hostname()
	if hostName != "" {
		*domain = hostName
	}
}

func validateInput(input *string) error {
	if *input == "" {
		return errors.New("❌ Input cannot be empty")
	}

	*input = strings.TrimSpace(*input)
	return nil
}
