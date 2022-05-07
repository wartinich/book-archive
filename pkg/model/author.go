package model

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var Authors = []Author{
	{Id: "1", Name: "Max"},
	{Id: "2", Name: "Arhip"},
}
