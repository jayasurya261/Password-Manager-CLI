package passwordManager

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"passwordManager/internal/db"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a credential by ID",
	Args:  cobra.ExactArgs(1), // require exactly 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Convert argument (string) to int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("invalid ID: %v", err)
		}

		err = db.DeleteCredential(id)
		if err != nil {
			log.Fatalf("failed to delete credential: %v", err)
		}

		fmt.Printf("Credential with ID %d deleted successfully 🗑️\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
