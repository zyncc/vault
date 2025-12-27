package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Secure Local Password Manager",
	Long: `The Best Password Manager which is fully secure and completely local, uses the best password hashing algorithms, also has some extra functionalities`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var uuidCmd = &cobra.Command{
	Use: "gen uuid",
	Short: "Generates a UUID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uuid.NewString())
	},
}

var genPassword = &cobra.Command{
	Use: "gen password",
	Short: "Generate Random Secure Password",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(uuidCmd)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vault.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


