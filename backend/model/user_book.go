package model

import "database/sql"

type UserBookModel interface {
	InsertUserBook(record *UserBook) error
	SelectUserBooks()
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
	return nil
}

func (m *userBookModel) SelectUserBooks() {

}

func (m *userBookModel) DeleteUserBook() {

}
