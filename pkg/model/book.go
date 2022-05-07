package model

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var Books = []Book{
	{Id: "1", Title: "Harry Potter"},
}
