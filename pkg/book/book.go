package book

import "library/pkg/author"

type Book struct {
	Key          string          `json:"key"`
	Title        string          `json:"title"`
	Authors      []author.Author `json:"authors"`
	EditionCount int64           `json:"edition_count"`
}
