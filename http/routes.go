package http

import (
	"regexp"

	"github.com/MarekWojt/GoMan/orm"
	"github.com/valyala/fasthttp"
)

type route struct {
	Path   *regexp.Regexp
	Action func(ctx *fasthttp.RequestCtx, parameters []string, session *orm.Session)
}

var routes []route = []route{
	{
		regexp.MustCompile("/api(/.*)"),
		handleAPIRequest,
	},
}
