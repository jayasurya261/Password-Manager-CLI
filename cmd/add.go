package passwordManager
import (
	"fmt"
	"github.com/spf13/cobra"
)

var service string
var username string
var password string
var addCmd = &cobra.Command{
	Use:"add",
	Short:"Add a new credential",
	Run:func(cmd *cobra.Command,args [] string){
		fmt.Printf("Added service=%s,Username=%s, Password=%s\n",service,username,password)
	},
	
	
}
func init(){
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&service,"service","s","","Service name (required)")
	addCmd.Flags().StringVarP(&username,"username","u","","Username (required)")
	addCmd.Flags().StringVarP(&password,"password","p","","Password (required)")

}