package main

import (
	"fmt"
	"postfix/mail"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")

	httprouter.Get("/postfix", func(ctx *httprouter.Context) {
		fmt.Fprintf(ctx.Ctx, "postfix api")
	})

	httprouter.Post("/send/text", mail.SendText)
	httprouter.Post("/send/html", mail.SendHTML)

	fhServer := &fasthttp.Server{
		Name:               "calicomoomoo mail service",
		Handler:            httprouter.FasthttpHandler(),
		MaxRequestBodySize: 5 * 1024 * 1024 * 1024 * 1024, // 5 TB
	}

	fhServer.ListenAndServe(":26")
}
