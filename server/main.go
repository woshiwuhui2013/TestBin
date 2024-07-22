package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	shared.cli = gin.New()
	shared.cli.Use(gin.Logger(), gin.Recovery())

}
