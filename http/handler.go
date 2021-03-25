package http

import (
	"fmt"
	"time"

	"github.com/MarekWojt/GoMan/orm"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

var parser fastjson.Parser = fastjson.Parser{}

func handleRequest(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	token := string(ctx.Request.Header.Cookie("SECURE_TOKEN"))

	key, session, err := orm.Sessions.GetSession(token)
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Println(err)
		return
	}

	newToken := fasthttp.Cookie{}
	newToken.SetKey("SECURE_TOKEN")
	newToken.SetValue(key)
	newToken.SetExpire(time.Now().Add(24 * time.Hour))
	newToken.SetHTTPOnly(true)
	newToken.SetSecure(true)
	ctx.Response.Header.SetCookie(&newToken)

	for _, route := range routes {
		if parameters := route.Path.FindStringSubmatch(path); len(parameters) > 0 {
			route.Action(ctx, parameters, session)
			return
		}
	}

	ctx.SetStatusCode(404)
}
