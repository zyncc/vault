package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Secure Local Password Manager",
	Long:  `The Best Password Manager which is fully secure and completely local, uses the best password hashing algorithms, also has some extra functionalities`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
