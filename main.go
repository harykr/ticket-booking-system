package main

import (
	"flag"

	"github.com/harykr/ticket-booking-system/app"
	"github.com/harykr/ticket-booking-system/db"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")
	port := flag.String("port", ":8080", "port")
	flag.Parse()

	if *migrate {
		db.RunMigrations()
		return
	}
	app.StartServer(*port)
}
