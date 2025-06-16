package db

import (
	"clinic-management/pkg/model"
)

func CreateUser(data model.User) error {

	query := `INSERT INTO users(id,email,password,role) VALUES ($1,$2,$3,$4)`

	_, err := conn.Exec(query, data.Id, data.Email, data.Password, data.Role)
	if err != nil {
		return err
	}

	return nil
}
func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := `SELECT id, email, password, role FROM users WHERE email=$1`
	err := conn.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}
