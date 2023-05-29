// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID    string  `json:"id"`
	Books []*Book `json:"books,omitempty"`
}

func (Author) IsEntity() {}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Author *Author `json:"author"`
}

func (Book) IsEntity() {}

type NewBook struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	AuthorID string `json:"authorId"`
}