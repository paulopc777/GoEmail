package main

import (
	handle "Email/server/Routes/Handle"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host  string
	Port  string
	Email string
	Pass  string
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("err open .env")
	}
	var environments = &Config{
		Host:  os.Getenv("HOST"),
		Port:  os.Getenv("PORT"),
		Email: os.Getenv("EMAIL"),
		Pass:  os.Getenv("PASS"),
	}

	app := gin.Default()

	v1 := app.Group("/api")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "ola mundo !"})
		})
		v1.POST("/send", handle.HandleSend(environments.Email, environments.Pass))
		v1.GET("/list-box", handle.HandleBox(environments.Email, environments.Pass))
		v1.GET("/list/:Inbox ", handle.EmailHandle(environments.Email, environments.Pass))
	}

	app.Run(":8080")
}
