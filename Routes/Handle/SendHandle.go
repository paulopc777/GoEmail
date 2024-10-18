package handle

import (
	client "Email/server/Client"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleSend(e string, p string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data client.Send_email
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"msg": "Error not data {to:string,title:string,body:string}"})
			return
		}
		fmt.Println(e)
		fmt.Println(p)
		s := client.Send(&data, e, p)
		if s {
			c.JSON(200, gin.H{"status": "Email as Send to " + data.To})
			return
		} else {
			c.JSON(404, gin.H{"status": "Not send email " + data.To})
			return
		}
	}

}
