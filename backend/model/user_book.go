package model

import (
	"database/sql"
	"fmt"
)

type UserBookModel interface {
	InsertUserBook(record *UserBook) error
	SelectUserBooksByUserID(userID string) ([]*UserBook, error)
	DeleteUserBook()
}

type userBookModel struct {
	DB *sql.DB
}

func NewUserBookModel(db *sql.DB) UserBookModel {
	return &userBookModel{DB: db}
}

type UserBook struct {
	ID       string
	UserID   string
	English  string
	Japanese string
}

func (m *userBookModel) InsertUserBook(record *UserBook) error {
	stmt, err := m.DB.Prepare("INSERT INTO user_book(id, user_id, eng, ja) VALUES(?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("model.InsertUserBook: %w", err)
	}
	_, err = stmt.Exec(record.ID, record.UserID, record.English, record.Japanese)
	if err != nil {
		return fmt.Errorf("model.InsertUserBook: %w", err)
	}
	return nil
}

func (m *userBookModel) SelectUserBooksByUserID(userID string) ([]*UserBook, error) {
	rows, err := m.DB.Query("SELECT * FROM user_book WHERE user_id=?", userID)
	if err != nil {
		return nil, fmt.Errorf("model.SelectUserBooks: %w", err)
	}
	return convertToUserBooks(rows)
}

func (m *userBookModel) DeleteUserBook() {

}

func convertToUserBooks(rows *sql.Rows) ([]*UserBook, error) {
	var records []*UserBook
	for rows.Next() {
		record := UserBook{}
		if err := rows.Scan(&record.ID, &record.UserID, &record.English, &record.Japanese); err != nil {
			return nil, fmt.Errorf("model.convertToUserBooks: %w", err)
		}
		records = append(records, &record)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("model.convertToUserBooks: %w", err)
	}
	return records, nil
}

//func convertToUserBook(row *sql.Row) (*UserBook, error) {
//	record := UserBook{}
//	err := row.Scan(&record.ID, &record.UserID, &record.English, &record.Japanese)
//	if err != nil {
//		return nil, fmt.Errorf("model.convertToUserBook: %w", err)
//	}
//	return &record, nil
//}
