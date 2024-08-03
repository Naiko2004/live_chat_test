package main

import (
	"github.com/gin-gonic/gin"

	"backend_chat/db"
	"backend_chat/router"
)

func main() {

	db.StartDbEngine()

	r := gin.Default()
	router.MapUrls(r)
	r.Run(":8080")
}
