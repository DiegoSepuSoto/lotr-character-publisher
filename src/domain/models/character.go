package models

type CharactersFile struct {
	Characters []*Character `json:"characters"`
}

type Character struct {
	Link     string `json:"link"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Category string `json:"category"`
}
