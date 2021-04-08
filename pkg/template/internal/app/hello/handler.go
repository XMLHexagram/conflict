package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello go-conflict")
}
