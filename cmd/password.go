package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate Random Secure Password",
	Run: func(cmd *cobra.Command, args []string) {

		password, err := generatePassword(length)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(password)
	},
}

var length int

func init() {
	passwordCmd.Flags().IntVarP(&length, "length", "l", 10, "specify the length of the password")
	genCmd.AddCommand(passwordCmd)
}

func generatePassword(length int) (string, error) {
	if length < 10 {
		return "", fmt.Errorf("length must be atleast 10 characters long")
	} else if length > 32 {
		return "", fmt.Errorf("length is too long")
	}

	const (
		letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits   = "0123456789"
		specials = "!@#$%^&*"
		allChars = letters + digits + specials
	)

	password := make([]byte, length)

	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
	if err != nil {
		return "", fmt.Errorf("failed to generate digit: %v", err)
	}
	password[0] = digits[num.Int64()]

	spec, err := rand.Int(rand.Reader, big.NewInt(int64(len(specials))))
	if err != nil {
		return "", fmt.Errorf("failed to generate special: %v", err)
	}
	password[1] = specials[spec.Int64()]

	for i := 2; i < length; i++ {
		charIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", fmt.Errorf("something went wrong, try again")
		}
		password[i] = allChars[charIdx.Int64()]
	}

	for i := length - 1; i > 0; i-- {
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", fmt.Errorf("failed to shuffle")
		}
		j := jBig.Int64()
		password[i], password[j] = password[j], password[i]
	}

	return string(password), nil
}
