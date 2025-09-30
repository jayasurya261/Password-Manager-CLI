package passwordManager

import (
	"fmt"
	"log"
	"os"

	"passwordManager/internal/crypto"
	"passwordManager/internal/db"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new credential (encrypted)",
	Run: func(cmd *cobra.Command, args []string) {
		service, _ := cmd.Flags().GetString("service")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if service == "" || username == "" || password == "" {
			log.Fatal("service, username and password are required")
		}

		// prompt master password (no echo)
		fmt.Print("Enter master password: ")
		mpBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			log.Fatalf("failed to read master password: %v", err)
		}
		master := string(mpBytes)

		// encrypt
		cipherB64, saltB64, nonceB64, err := crypto.Encrypt([]byte(password), master)
		if err != nil {
			log.Fatalf("encryption failed: %v", err)
		}

		// save encrypted fields
		if err := db.AddCredential(service, username, cipherB64, saltB64, nonceB64); err != nil {
			log.Fatalf("failed to save credential: %v", err)
		}
		fmt.Println("Credential saved successfully ✅")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("service", "s", "", "Service name (e.g., Gmail, GitHub)")
	addCmd.Flags().StringP("username", "u", "", "Username (required)")
	addCmd.Flags().StringP("password", "p", "", "Password (required)")
}
