package domain

import "time"

type Book struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorID    Author    `json:"author_id"`
	PublishDate time.Time `json:"publish_date"`
}

var Books = []Book{
	{Id: "1", Title: "Harry Potter", Description: "Amazing book", AuthorID: Author{Id: "1", Name: "Max"}, PublishDate: time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)},
}
