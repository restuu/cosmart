package webserver

import (
	bookRest "library/pkg/book/http/rest"
	bookRepository "library/pkg/book/repository"
	bookService "library/pkg/book/service"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	bookRepo := bookRepository.NewBookRepository()
	pickupRepo := bookRepository.NewBookScheduleRepository()

	bookGettingService := bookService.NewBookGettingService(bookRepo)
	bookPickupAddingService := bookService.NewBookScheduleAddingService(pickupRepo)
	bookScheduleGettingService := bookService.NewBookScheduleGettingService(pickupRepo)

	bookRest.Router(r, bookGettingService, bookPickupAddingService, bookScheduleGettingService)
}
