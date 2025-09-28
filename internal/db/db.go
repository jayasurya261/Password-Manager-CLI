package db
import (
	"database/sql"
	"fmt"
	 _ "modernc.org/sqlite"
)
var DB *sql.DB
func InitDB(filepath string)error{
	var err error
	DB,err = sql.Open("sqlite",filepath)
	if err != nil{
		return err
	}
	createTable := `CREATE TABLE IF NOT EXISTS credentials (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
);`
	_,err = DB.Exec(createTable)
	if err != nil{
		return fmt.Errorf("failed to create table: %v",err)
	}
	return nil
}
func GetAllCredentials()([]map[string]string, error){
	rows, err := DB.Query("SELECT id, username, password FROM credentials")
	if err !=nil{
		return nil,err
	}
	defer rows.Close()
	var credentials []map[string]string
	for rows.Next(){
		var id int
		var username,password string
		if err := rows.Scan(&id,&username,&password); err != nil{
			return nil,err
		}
		credentials = append(credentials,map[string]string{
			"id": fmt.Sprintf("%d",id),
			"username": username,
			"password": password,
		})
		
	}
	
return credentials,nil

}