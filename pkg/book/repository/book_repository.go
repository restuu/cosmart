package repository

import (
    "encoding/json"
    "fmt"
    "library/pkg/book"
    "net/http"
)

type BookRepository interface {
    FindByGenre(genre string) ([]book.Book, error)
}

type bookRepository struct {

}

const (
    baseUrl string = "http://openlibrary.org/subjects"
)

type BookApiResponse struct {
    Works []book.Book `json:"works"`
}

func (br *bookRepository) FindByGenre(genre string) ([]book.Book, error) {
    fullUrl := baseUrl + fmt.Sprintf("/%s.json", genre)
    res, err := http.Get(fullUrl)

    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    decoder := json.NewDecoder(res.Body)

    bookResponse := BookApiResponse{}

    if err = decoder.Decode(&bookResponse); err != nil {
        return nil, err
    }

    return bookResponse.Works, nil
}

var _bookRepository BookRepository

func NewBookRepository() BookRepository {
    if _bookRepository == nil {
        _bookRepository = &bookRepository{}
    }

    return _bookRepository
}

