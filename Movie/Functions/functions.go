package Functions

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckERR(err error, str string) error {
	if err != nil {
		log.Println("loc:", str, "error:", err.Error())
	}
	return err
}

func CheckERROR(err error, str string) {
	if err != nil {
		log.Println("loc:", str, "error:", err.Error())
		return
	}
}

func CheckSERVERErr(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"ok":    false,
		"error": err,
	})
	return
}

func SERVEROKDATA(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
	return
}

func SERVEROKAY(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
	return
}

//
//func CheckSERVERBodyErr(err error,a int, c *gin.Context) {
//	c.JSON(http.StatusInternalServerError, gin.H{
//		"ok":    false,
//		"error": err,
//	})
//}
