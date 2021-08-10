package main

import (
	"flag"
	"fmt"

	"github.com/harykr/ticket-booking-system/app"
	"github.com/harykr/ticket-booking-system/db"
)

func main() {
	migrate := flag.Bool("migrate", false, "To run migrate")

	seed := flag.Bool("seed", false, "To run db seeding")
	flag.Parse()

	if *migrate {
		db.RunMigrations()
		return
	}

	if *seed {
		fmt.Println("running seeds")
	}

	app.StartServer()
}
