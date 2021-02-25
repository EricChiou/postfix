package mail

import (
	"encoding/json"
	"net/smtp"
	"postfix/apihandler"
	req "postfix/vo/request"

	"github.com/EricChiou/httprouter"
)

// SendText email
func SendText(ctx *httprouter.Context) {
	var reqVo req.SendMailData
	if err := json.Unmarshal(ctx.Ctx.PostBody(), &reqVo); err != nil {
		apihandler.Error(ctx, err.Error())
		return
	}

	data, err := getSendData(reqVo)
	if err != nil {
		apihandler.Error(ctx, err.Error())
		return
	}

	header := make(map[string]string)
	header["From"] = data.FromHeader
	header["To"] = data.ToHeader
	header["Subject"] = data.Subject
	msg := getMsg(header, reqVo.Body)

	err = send(data, msg)
	if err != nil {
		apihandler.Error(ctx, err.Error())
		return
	}

	apihandler.OK(ctx, "send text email")
}

// SendHTML email
func SendHTML(ctx *httprouter.Context) {
	var reqVo req.SendMailData
	if err := json.Unmarshal(ctx.Ctx.PostBody(), &reqVo); err != nil {
		apihandler.Error(ctx, err.Error())
		return
	}

	data, err := getSendData(reqVo)
	if err != nil {
		apihandler.Error(ctx, err.Error())
		return
	}

	header := make(map[string]string)
	header["From"] = data.FromHeader
	header["To"] = data.ToHeader
	header["Subject"] = data.Subject
	header["Content-Type"] = `text/html; charset="UTF-8"`
	msg := getMsg(header, reqVo.Body)

	err = send(data, msg)
	if err != nil {
		apihandler.Error(ctx, err.Error())
		return
	}

	apihandler.OK(ctx, "send text email")
}

func combineMsg() {

}

func send(data sendData, msg string) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}

	if err = c.Mail(data.FromEmail.Email); err != nil {
		return err
	}

	for _, emailData := range data.ToEmails {
		if err = c.Rcpt(emailData.Email); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = c.Quit()
	if err != nil {
		return err
	}

	return nil
}
