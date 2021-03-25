package http

import (
	encodingJson "encoding/json"

	"github.com/MarekWojt/GoMan/orm"
	"github.com/MarekWojt/GoMan/util/json"
	"github.com/MarekWojt/GoMan/util/response"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

func handleAPIRequest(ctx *fasthttp.RequestCtx, parameters []string, session *orm.Session) {
	path := parameters[0]
	body := ctx.Request.Body()

	var data *fastjson.Value
	var err error

	if len(body) > 0 {
		data, err = parser.ParseBytes(body)
		if err != nil {
			ctx.SetStatusCode(400)
			return
		}
	} else {
		data = fastjson.MustParse("null")
	}

	for _, route := range apiRoutes {
		if parameters := route.Path.FindStringSubmatch(path); len(parameters) > 0 {
			var apiResponse response.Response
			if route.HasAccess(session) {
				apiResponse = route.Action(json.JSON{Value: data}, parameters, session)
			} else {
				apiResponse = response.New(403, "You do not have access to this action")
			}

			if apiResponse.Status == 0 {
				apiResponse.Status = 200
			}

			ctx.SetStatusCode(apiResponse.Status)
			parsed, err := encodingJson.Marshal(apiResponse)
			if err != nil {
				parsed, err = encodingJson.Marshal(response.New(500, "internal server error"))
				if err != nil {
					println(err.Error())
					return
				}
			}

			ctx.SetBody(parsed)

			return
		}
	}

	ctx.SetStatusCode(404)
}
