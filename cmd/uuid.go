package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates a UUID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uuid.NewString())
	},
}

func init() {
	genCmd.AddCommand(uuidCmd)
}
