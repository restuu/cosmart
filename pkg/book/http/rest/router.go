package rest

import (
	"library/pkg/book/model"
	"library/pkg/book/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookGettingService service.BookGettingService
}

// Router ...
func Router(
    r *gin.Engine,
    bookGettingService service.BookGettingService,
    ) {
	handler := &bookHandler{
        bookGettingService: bookGettingService,
    }

	books := r.Group("/books")

	books.GET("/", handler.findBooks)
}

func (h *bookHandler) findBooks(c *gin.Context) {

	bookFindRequest := model.BookFindRequest{}

	if err := c.ShouldBindQuery(&bookFindRequest); err != nil {
		handleInvalidParam(c)
		return
	}

	books, err := h.bookGettingService.FindBooks(bookFindRequest)
	if err != nil {
		handleInvalidParam(c)
		return
	}

	c.JSON(http.StatusOK, books)
}

func handleInvalidParam(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "request failed"})
}
