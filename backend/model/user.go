package model

import (
	"database/sql"
	"fmt"
)

type UserModel interface {
	InsertUser(record *User) error
	SelectUserByPrimaryKey(userID string) (*User, error)
	SelectUserByMail(mail string) (*User, error)
	UpdateUserByPrimaryKey(record *User) error
}

type userModel struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) UserModel {
	return &userModel{
		DB: db,
	}
}

// User user data
type User struct {
	UserID   string
	Mail     string
	Password string
}

// InsertUser insert user
func (m *userModel) InsertUser(record *User) error {
	stmt, err := m.DB.Prepare("INSERT INTO user(id, mail, password) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("model.InsertUser: %w", err)
	}
	_, err = stmt.Exec(record.UserID, record.Mail, record.Password)
	if err != nil {
		return fmt.Errorf("model.InsertUser: %w", err)
	}
	return nil
}

// SelectUserByPrimaryKey get user
func (m *userModel) SelectUserByPrimaryKey(userID string) (*User, error) {
	row := m.DB.QueryRow("SELECT * FROM user WHERE id=?", userID)
	return convertToUser(row)
}

// SelectUserByMail get user
func (m *userModel) SelectUserByMail(mail string) (*User, error) {
	row := m.DB.QueryRow("SELECT * FROM user WHERE mail=?", mail)
	return convertToUser(row)
}

// UpdateUserByPrimaryKey update user
func (m *userModel) UpdateUserByPrimaryKey(record *User) error {
	stmt, err := m.DB.Prepare("UPDATE user SET mail = ?, password = ? WHERE id = ? ")
	if err != nil {
		return fmt.Errorf("model.UpdateUserByPrimaryKey: %w", err)
	}
	_, err = stmt.Exec(record.Mail, record.Password, record.UserID)
	if err != nil {
		return fmt.Errorf("model.UpdateUserByPrimaryKey: %w", err)
	}
	return nil
}

// ConvertToUser convert row to UserDTO
func convertToUser(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.UserID, &user.Mail, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("model.convertToUser: %w", err)
	}
	return &user, nil
}
