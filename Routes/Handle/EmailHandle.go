package handle

import (
	client "Email/server/Client"

	"github.com/gin-gonic/gin"
)

func EmailHandle(e string, p string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var imap client.ImapClient
		imap.LoginImap(e, p)
		res := imap.GetEmails(e, p)

		c.JSON(200, gin.H{"title": res})

	}
}
