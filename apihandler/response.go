package apihandler

import (
	"encoding/json"
	"fmt"
	res "postfix/vo/response"

	"github.com/EricChiou/httprouter"
)

// OK reponse
func OK(ctx *httprouter.Context, data interface{}) {
	resp := res.Response{
		Status: "ok",
		Data:   data,
	}

	bytes, _ := json.Marshal(resp)
	fmt.Fprintf(ctx.Ctx, string(bytes))
}

// Error reponse
func Error(ctx *httprouter.Context, err string) {
	resp := res.Response{
		Status: "error",
		Trace:  err,
	}

	bytes, _ := json.Marshal(resp)
	fmt.Fprintf(ctx.Ctx, string(bytes))
}
