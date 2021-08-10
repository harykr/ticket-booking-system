package contract

import "github.com/harykr/ticket-booking-system/models"

type BookingResponse struct {
	Success bool            `json:"success"`
	Error   []string        `json:"errors"`
	Data    []models.Ticket `json:"data"`
}

type BookingRequest struct {
	Catalog models.Catalog `json:"catalog"`
	Slot    models.Slot    `json:"slot"`
}
