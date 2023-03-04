package domain

type Book struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type CreateBookInput struct{
	Title 		string 	  `json:"title"`
	Description string    `json:"description"`
}

var Books = []Book{
	{Id: "1", Title: "Harry Potter", Description: "Amazing book"},
}
