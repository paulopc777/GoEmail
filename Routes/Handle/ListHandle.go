package handle

import (
	client "Email/server/Client"

	"github.com/gin-gonic/gin"
)

func HandleBox(e string, p string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Init Imap
		var imap client.ImapClient
		imap.LoginImap(e, p)
		res := imap.ListEmailBox()

		ctx.JSON(200, gin.H{"inbox": res, "count": len(res)})
	}
}
