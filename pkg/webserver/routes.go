package webserver

import (
	bookRest "library/pkg/book/http/rest"
	bookRepository "library/pkg/book/repository"
	bookService "library/pkg/book/service"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	bookRepo := bookRepository.NewBookRepository()

	bookGettingService := bookService.NewBookGettingService(bookRepo)

	bookRest.Router(r, bookGettingService)
}
