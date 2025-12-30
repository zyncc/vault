package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zyncc/vault/db"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the vault to show all passwords",
	Run: func(cmd *cobra.Command, args []string) {
		// Verify Master Password
		if master != "Master" {
			fmt.Println("❌ Master Password does not match")
			return
		}

		// Fetch all passwords
		passwords, err := fetchAllPasswords()
		if err != nil {

		}
		for i := range passwords {
			fmt.Printf("%v          %v\n", passwords[i].Domain, passwords[i].Password)
		}
	},
}

var master string

func init() {
	openCmd.Flags().StringVarP(&master, "master", "m", "", "Master Password is required for accessing sensitive information")
	if err := openCmd.MarkFlagRequired("master"); err != nil {
		panic("master flag is required")
	}

	rootCmd.AddCommand(openCmd)
}

func fetchAllPasswords() ([]db.PasswordStore, error) {
	q := db.Init()

	return q.GetAllPasswords(context.Background())
}
