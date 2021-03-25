package http

import "github.com/MarekWojt/GoMan/orm"

func accessAlways(session *orm.Session) bool {
	return true
}

func accessLoggedIn(session *orm.Session) bool {
	return session.User.ID != 0
}

func accessAdmin(session *orm.Session) bool {
	return session.User.IsAdmin
}
