package mail

import (
	"errors"
	"fmt"
	"net/mail"
	req "postfix/vo/request"
	"strings"
)

type sendData struct {
	FromHeader string
	ToHeader   string
	ToEmails   []req.EmailData
	Subject    string
}

func getSendData(mailData req.SendMailData) (sendData, error) {
	var h sendData

	fromHeader, err := getFromHeader(mailData.From)
	if err != nil {
		return h, err
	}
	h.FromHeader = fromHeader

	toHeader, err := getToHeader(mailData.To)
	if err != nil {
		return h, err
	}
	h.ToHeader = toHeader

	if len(mailData.Subject) == 0 {
		return h, errors.New("subject can not be empty")
	}
	h.Subject = mailData.Subject

	return h, nil
}

func getFromHeader(from req.EmailData) (string, error) {
	if len(from.Name) == 0 || len(from.Email) == 0 {
		return "", errors.New("from email data can not be empty")
	}

	addr := mail.Address{Name: from.Name, Address: from.Email}
	return addr.String(), nil
}

func getToHeader(to []req.EmailData) (string, error) {
	var toAddrs []string
	for _, data := range to {
		if len(data.Name) == 0 || len(data.Email) == 0 {
			return "", errors.New("to email data can not be empty")
		}

		addr := mail.Address{Name: data.Name, Address: data.Email}
		toAddrs = append(toAddrs, addr.String())
	}

	return strings.Join(toAddrs, ", "), nil
}

func getMsg(header map[string]string, body string) string {
	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body

	return msg
}
