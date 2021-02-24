package mail

import (
	"fmt"
	"postfix/apihandler"

	"github.com/EricChiou/httprouter"
)

// SendText email
func SendText(ctx *httprouter.Context) {
	// send text email

	apihandler.OK(ctx, "send text email")
}

// SendHTML email
func SendHTML(ctx *httprouter.Context) {
	// send html email

	fmt.Fprintf(ctx.Ctx, "send html email")
}
