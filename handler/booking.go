package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harykr/ticket-booking-system/contract"
	"github.com/harykr/ticket-booking-system/models"
	"github.com/satori/uuid"
)

func BookTicket() func(c *gin.Context) {
	return func(c *gin.Context) {
		var bookingRequest contract.BookingRequest

		if err := c.ShouldBindJSON(&bookingRequest); err != nil {
			c.JSON(http.StatusCreated, contract.BookingResponse{
				Success: false,
				Errors:  []string{"Invalid request"},
				Data:    nil,
			})
			return
		}

		ticketId := uuid.NewV4()
		ticket := models.Ticket{
			Id:      ticketId,
			Catalog: bookingRequest.Catalog,
			Slot:    bookingRequest.Slot,
		}

		c.JSON(http.StatusCreated, contract.BookingResponse{
			Success: true,
			Errors:  []string{},
			Data: []models.Ticket{
				ticket,
			},
		})

	}
}
