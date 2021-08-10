package app

func StartServer(port string) {

	r := CreateRouter()
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
