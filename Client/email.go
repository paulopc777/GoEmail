package client

import (
	entity "Email/server/Entity"
	"fmt"
	"net/smtp"
	"os"
)

func Send(data *entity.Email, user string, pass string) bool {
	// ENV
	smtpHost := os.Getenv("HOST")
	smtpPort := os.Getenv("PORT")
	to := []string{data.To}
	message := []byte("Subject:" + data.Title + "\n" +
		"\n" + data.Body)

	auth := smtp.PlainAuth("", user, pass, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, user, to, message)
	if err != nil {
		fmt.Println("Erro ao enviar o e-mail:", err)
		return false
	}
	fmt.Println("E-mail enviado com sucesso!")
	return true
}
