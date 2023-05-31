package app

import (
	"authentication/routes"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()
	r = routes.AuthenticationRoutes(r)
	
	r.Run()
}