package entity

type Email struct {
	To      string `form:to binding:"required`
	Subject string
	From    string
	Title   string `form:title binding:"required,email"`
	Body    string `form:body `
}

func (e *Email) CreateNewEmail(Send Email) Email {

	var NewEmail Email

	NewEmail.To = Send.To
	NewEmail.Title = Send.Title
	NewEmail.Body = Send.Body

	return NewEmail
}
