package service

import (
	"fmt"
	"library/pkg/book"
	"library/pkg/book/model"
	"library/pkg/book/repository"
)

type BookGettingService interface {
	FindBooks(request model.BookFindRequest) ([]book.Book, error)
}

type bookGettingService struct {
	bookRepository repository.BookRepository
}

func (b bookGettingService) FindBooks(request model.BookFindRequest) ([]book.Book, error) {
	// only support if one of the parameter fulfilled
	if request.Genre == "" {
		return nil, fmt.Errorf("missing search parameter")
	}

	return b.bookRepository.FindByGenre(request.Genre)
}

var _bookGettingService BookGettingService

func NewBookGettingService(
	bookRepository repository.BookRepository,
) BookGettingService {

	if _bookGettingService == nil {
		_bookGettingService = &bookGettingService{
			bookRepository: bookRepository,
		}
	}

	return _bookGettingService
}