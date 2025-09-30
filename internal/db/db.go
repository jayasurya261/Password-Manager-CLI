package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes database and creates table with cipher columns
func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite", filepath)
	if err != nil {
		return err
	}

	createTable := `CREATE TABLE IF NOT EXISTS credentials (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		service TEXT NOT NULL,
		username TEXT NOT NULL,
		password_cipher TEXT NOT NULL,
		salt TEXT NOT NULL,
		nonce TEXT NOT NULL
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	return nil
}

// AddCredential stores encrypted credential fields into DB
func AddCredential(service, username, passwordCipher, salt, nonce string) error {
	stmt := `INSERT INTO credentials(service, username, password_cipher, salt, nonce) VALUES(?, ?, ?, ?, ?)`
	_, err := DB.Exec(stmt, service, username, passwordCipher, salt, nonce)
	return err
}

// GetAllCredentials returns rows with cipher data
func GetAllCredentials() ([]map[string]string, error) {
	rows, err := DB.Query("SELECT id, service, username, password_cipher, salt, nonce FROM credentials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []map[string]string
	for rows.Next() {
		var id int
		var service, username, passwordCipher, salt, nonce string
		if err := rows.Scan(&id, &service, &username, &passwordCipher, &salt, &nonce); err != nil {
			return nil, err
		}
		credentials = append(credentials, map[string]string{
			"id":              fmt.Sprintf("%d", id),
			"service":         service,
			"username":        username,
			"password_cipher": passwordCipher,
			"salt":            salt,
			"nonce":           nonce,
		})
	}
	return credentials, nil
}

// DeleteCredential deletes by ID
func DeleteCredential(id int) error {
	stmt, err := DB.Prepare("DELETE FROM credentials WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
