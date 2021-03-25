package http

import (
	"regexp"

	"github.com/MarekWojt/GoMan/api"
	"github.com/MarekWojt/GoMan/orm"
	"github.com/MarekWojt/GoMan/util/json"
	"github.com/MarekWojt/GoMan/util/response"
)

type apiRoute struct {
	Path      *regexp.Regexp
	Action    func(data json.JSON, parameters []string, session *orm.Session) response.Response
	HasAccess func(session *orm.Session) bool
}

var apiRoutes []apiRoute = []apiRoute{
	{
		regexp.MustCompile("/login"),
		api.UserLogin,
		accessAlways,
	},
	{
		regexp.MustCompile("/register"),
		api.UserRegister,
		accessAlways,
	},
	{
		regexp.MustCompile("/whoami"),
		api.UserWhoAmI,
		accessAlways,
	},
	{
		regexp.MustCompile("/logout"),
		api.UserLogout,
		accessLoggedIn,
	},
}
