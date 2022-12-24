package repository

import "library/pkg/book/model"

type BookScheduleRepository interface {
	FindAll() ([]model.BookSchedule, error)
	AddSchedule(model.BookSchedule) error
}

type bookScheduleRepository struct {
	schedules []model.BookSchedule
}

func (r *bookScheduleRepository) FindAll() ([]model.BookSchedule, error) {
	return r.schedules, nil
}

func (r *bookScheduleRepository) AddSchedule(schedule model.BookSchedule) error {
	r.schedules = append(r.schedules, schedule)

	return nil
}

func NewBookScheduleRepository() BookScheduleRepository {
	return &bookScheduleRepository{
		schedules: make([]model.BookSchedule, 0),
	}
}
