package service

import (
	"library/pkg/book/model"
	"library/pkg/book/repository"
)

// BookScheduleGettingService ...
type BookScheduleGettingService interface {
	FindAll() ([]model.BookSchedule, error)
}

type bookScheduleGettingService struct {
	bookScheduleRepository repository.BookScheduleRepository
}

func (s *bookScheduleGettingService) FindAll() ([]model.BookSchedule, error) {
	return s.bookScheduleRepository.FindAll()
}

var _bookScheduleGettingService BookScheduleGettingService

func NewBookScheduleGettingService(
	bookScheduleRepository repository.BookScheduleRepository,
) BookScheduleGettingService {

	if _bookScheduleGettingService == nil {
		_bookScheduleGettingService = &bookScheduleGettingService{
			bookScheduleRepository: bookScheduleRepository,
		}
	}

	return _bookScheduleGettingService
}
