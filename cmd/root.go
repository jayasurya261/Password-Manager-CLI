package passwordManager
import(
	"fmt"
	"github.com/spf13/cobra"
	"passwordManager/internal/db"
	"log"
)

var rootCmd = &cobra.Command{
Use: "pwm",
Short : "Password Manager CLI",
Long: `A CLI password manager built with security and simplicity`,
Run :func(cmd *cobra.Command,args[] string){
fmt.Println("Welcome to PWM - Your Password Manager CLI")
},

}
func Execute() error{
	err:=db.InitDB("passwords.db")
	if err != nil{
		log.Fatalf("database initialization failed: %v",err)

	}
return rootCmd.Execute()
}