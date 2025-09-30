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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored credentials (prompts master password to decrypt)",
	Run: func(cmd *cobra.Command, args []string) {
		// prompt master password once
		fmt.Print("Enter master password: ")
		mpBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			log.Fatalf("failed to read master password: %v", err)
		}
		master := string(mpBytes)

		creds, err := db.GetAllCredentials()
		if err != nil {
			log.Fatalf("failed to fetch credentials: %v", err)
		}
		if len(creds) == 0 {
			fmt.Println("No credentials found.")
			return
		}

		fmt.Println("Stored credentials:")
		for _, c := range creds {
			plain := "<decryption failed>"
			pt, err := crypto.Decrypt(c["password_cipher"], c["salt"], c["nonce"], master)
			if err == nil {
				plain = string(pt)
			}
			fmt.Printf("ID: %s | Service: %s | Username: %s | Password: %s\n",
				c["id"], c["service"], c["username"], plain)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
