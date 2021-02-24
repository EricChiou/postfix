package apihandler

import (
	"encoding/json"
	"fmt"
	"postfix/vo"

	"github.com/EricChiou/httprouter"
)

// OK reponse
func OK(ctx *httprouter.Context, data interface{}) {
	resp := vo.Response{
		Status: "ok",
		Data:   data,
	}

	bytes, _ := json.Marshal(resp)
	fmt.Fprintf(ctx.Ctx, string(bytes))
}
