package webserver

import (
    bookRest "library/pkg/book/http/rest"
    bookService "library/pkg/book/service"
    bookRepository "library/pkg/book/repository"

    "github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {

    bookRepo := bookRepository.NewBookRepository()

    bookGettingService := bookService.NewBookGettingService(bookRepo)

    bookRest.Router(r, bookGettingService)
}
