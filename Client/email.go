package client

import (
	"fmt"
	"net/smtp"
	"os"
)

type Send_email struct {
	To    string `form:to binding:"required`
	Title string `form:title binding:"required,email"`
	Body  string `form:body `
}


func Send(data *Send_email, e string, p string) bool {
	// ENV
	smtpHost := os.Getenv("HOST")
	smtpPort := os.Getenv("PORT")
	to := []string{data.To}
	message := []byte("Subject:" + data.Title + "\n" +
		"\n" + data.Body)

	auth := smtp.PlainAuth("", e, p, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, e, to, message)
	if err != nil {
		fmt.Println("Erro ao enviar o e-mail:", err)
		return false
	}
	fmt.Println("E-mail enviado com sucesso!")
	return true
}
