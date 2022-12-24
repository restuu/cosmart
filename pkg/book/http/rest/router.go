package rest

import (
	"library/pkg/book/model"
	"library/pkg/book/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookGettingService         service.BookGettingService
	bookPickupAddingService    service.BookPickupAddingService
	bookScheduleGettingService service.BookScheduleGettingService
}

// Router ...
func Router(
	r *gin.Engine,
	bookGettingService service.BookGettingService,
	bookPickupAddingService service.BookPickupAddingService,
	bookScheduleGettingService service.BookScheduleGettingService,
) {

	handler := &bookHandler{
		bookGettingService:         bookGettingService,
		bookPickupAddingService:    bookPickupAddingService,
		bookScheduleGettingService: bookScheduleGettingService,
	}

	books := r.Group("/books")

	books.GET("/", handler.findBooks)
	books.POST("/pickup", handler.submitPickupSchedule)
	books.GET("/schedules", handler.findAllSchedules)
}

func (h *bookHandler) findBooks(c *gin.Context) {
	bookFindRequest := model.BookFindRequest{}

	if err := c.ShouldBindQuery(&bookFindRequest); err != nil {
		handleInvalidParam(c, err)
		return
	}

	books, err := h.bookGettingService.FindBooks(bookFindRequest)
	if err != nil {
		handleInvalidParam(c, err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *bookHandler) submitPickupSchedule(c *gin.Context) {
	bookPickupSchedule := model.BookSchedule{}

	if err := c.ShouldBindJSON(&bookPickupSchedule); err != nil {
		handleInvalidParam(c, err)
		return
	}

	if err := h.bookPickupAddingService.AddSchedule(bookPickupSchedule); err != nil {
		handleInvalidParam(c, err)
	}

	c.JSON(200, gin.H{"status": "OK"})
}

func (h *bookHandler) findAllSchedules(c *gin.Context) {
	res, err := h.bookScheduleGettingService.FindAll()

	if err != nil {
		handleInvalidParam(c, err)
		return
	}

	c.JSON(200, res)
}

func handleInvalidParam(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "request failed", "reason": err.Error()})
}
