package app

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func CreateRouter() *Router {
	r := &Router{gin.Default()}
	r.addRoutes()
	return r
}

func (r Router) addRoutes() {
	r.POST("/book", handler.bookTicket())
}
