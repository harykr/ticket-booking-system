package app

import (
	"fmt"
	"net/http"
)

func StartServer() {

	s := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", pingHandler())
	http.HandleFunc("/book", bookHandler())

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}
	fmt.Println("Server Started")
}

func pingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}

func bookHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Ticket Booked")
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte("Ticket Booked."))
	}
}
