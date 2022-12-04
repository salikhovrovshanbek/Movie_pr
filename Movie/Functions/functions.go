package Functions

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckERR(err error) {
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
}

func CheckSERVERErr(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"ok":    false,
		"error": err,
	})
}

func SERVEROKDATA(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func SERVEROKAY(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
