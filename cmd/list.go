package passwordManager
import (
	"fmt"
	"github.com/spf13/cobra"
	"passwordManager/internal/db"
	"log"
)
var listCmd = &cobra.Command{
	Use: "list",
	Short: "List all stored credentials",
	Run: func(cmd *cobra.Command, args []string) {
		credentials, err := db.GetAllCredentials()
		if err != nil {
			log.Fatalf("Failed to retrieve credentials: %v", err)
		}
		if len(credentials) == 0 {
			fmt.Println("No credentials found.")
			return
		}
		fmt.Println("Stored Credentials:")
		for _, cred := range credentials {
			fmt.Printf("ID: %s, Username: %s, Password: %s\n", cred["id"], cred["username"], cred["password"])
		}
	},

}
func init(){
	rootCmd.AddCommand(listCmd)
}