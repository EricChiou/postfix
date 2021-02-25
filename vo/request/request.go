package request

// SendMailData request vo
type SendMailData struct {
	From    EmailData   `json:"from"`
	To      []EmailData `json:"to"`
	Subject string      `json:"subject"`
	Body    string      `json:"body"`
}

// EmailData include Name and Email
type EmailData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
