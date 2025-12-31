package cmd

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/zyncc/vault/db"
)

var insertCmd = cobra.Command{
	Use:   "insert",
	Short: "Insert a new password into the vault",
	Run: func(cmd *cobra.Command, args []string) {
		if err := insertPasswordToDatabase(&domain, &password); err != nil {
			if strings.Contains(err.Error(), "2067") {
				fmt.Println("❌ An Existing Password exists for", domain)
				return
			}
			fmt.Println("❌ Failed to add password to vault")
			return
		}
		fmt.Printf("✅ Succesfully added password for %v to the vault\n", domain)
	},
}

var domain string
var password string

func init() {
	insertCmd.Flags().StringVarP(&domain, "domain", "d", "", "The Domain name of the website's password you want to insert. ex: reddit.com")
	insertCmd.Flags().StringVarP(&password, "password", "p", "", "The password that you want to insert")

	if err := insertCmd.MarkFlagRequired("domain"); err != nil {
		panic("something went wrong, try again")
	}
	if err := insertCmd.MarkFlagRequired("password"); err != nil {
		panic("something went wrong, try again")
	}

	rootCmd.AddCommand(&insertCmd)
}

func insertPasswordToDatabase(domain *string, password *string) error {
	parseDomain(domain)
	q := db.Init()
	return q.InsertIntoPasswordStore(context.Background(), db.InsertIntoPasswordStoreParams{
		ID:       uuid.NewString(),
		Domain:   *domain,
		Password: *password,
	})
}

func parseDomain(domain *string) {
	input := strings.TrimSpace(*domain)
	if input == "" {
		return
	}

	u, err := url.Parse(*domain)
	if err != nil {
		return 
	}

	*domain = strings.ToLower(u.Hostname())
}
