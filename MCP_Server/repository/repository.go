package repository

import (
	"database/sql"
	"mcp_sql_server_/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User

	query := "SELECT UserID,FirstName,LastName,Email FROM Users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetContactOfUsers(db *sql.DB) ([]map[string]string, error) {
	var results []map[string]string

	query := "Select Users.FirstName,UserContacts.PhoneNumber From Users Join UserContacts ON Users.UserID = UserContacts.UserID"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var firstName string
		var phone string
		err := rows.Scan(&firstName, &phone)
		if err != nil {
			return nil, err
		}
		rows := map[string]string{
			"FirstName": firstName,
			"Phone":     phone,
		}
		results = append(results, rows)
	}
	return results, nil
}
