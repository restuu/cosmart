package service

import (
	"library/pkg/book/model"
	"library/pkg/book/repository"
)

// BookPickupAddingService ...
type BookPickupAddingService interface {
	AddSchedule(model.BookSchedule) error
}

type bookPickupAddingService struct {
	bookPickupScheduleRepository repository.BookScheduleRepository
}

func (pas *bookPickupAddingService) AddSchedule(schedule model.BookSchedule) error {
	return pas.bookPickupScheduleRepository.AddSchedule(schedule)
}

var _bookPickupAddingService BookPickupAddingService

// NewBookScheduleAddingService ...
func NewBookScheduleAddingService(
	bookPickupScheduleRepository repository.BookScheduleRepository,
) BookPickupAddingService {

	if _bookPickupAddingService == nil {
		_bookPickupAddingService = &bookPickupAddingService{
			bookPickupScheduleRepository: bookPickupScheduleRepository,
		}
	}

	return _bookPickupAddingService
}
