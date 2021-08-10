package models

import "github.com/satori/uuid"

type Ticket struct {
	Id      uuid.UUID `json:"id"`
	Catalog Catalog   `json:"catalog"`
	Slot    Slot      `json:"slot"`
}
