package api

import (
	"github.com/MarekWojt/GoMan/db"
	"github.com/MarekWojt/GoMan/orm"
	"github.com/MarekWojt/GoMan/util/json"
	"github.com/MarekWojt/GoMan/util/response"
)

// UserLogin is the user login action
func UserLogin(data json.JSON, parameters []string, session *orm.Session) response.Response {
	username, err := data.GetString("username")
	if err != nil {
		return response.New(400, "invalid username")
	}

	password, err := data.GetString("password")
	if err != nil {
		return response.New(400, "invalid password")
	}

	user := orm.User{}
	db.DB.Model(&orm.User{}).Where(map[string]interface{}{"username": username}).Take(&user)

	authed, err := user.Auth(password)
	if err != nil {
		return response.New(500, err.Error())
	}

	if !authed {
		return response.New(401, "Invalid credentials or disabled account")
	}

	session.User = user
	return response.Ok(user)
}

// UserRegister is the register action
func UserRegister(data json.JSON, parameters []string, session *orm.Session) response.Response {
	email, err := data.GetString("email")
	if err != nil {
		return response.New(400, "invalid username")
	}

	username, err := data.GetString("username")
	if err != nil {
		return response.New(400, "invalid email")
	}

	password, err := data.GetString("password")
	if err != nil {
		return response.New(400, "invalid password")
	}

	user := orm.User{
		Username:      username,
		Email:         email,
		PlainPassword: password,
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		return response.New(409, "could not be created")
	}

	return response.Response{Data: user}
}

// UserLogout is the logout action
func UserLogout(data json.JSON, parameters []string, session *orm.Session) response.Response {
	session.User = orm.User{}
	return response.New(200, nil)
}

// UserWhoAmI is an action that returns the current user or null if logged out
func UserWhoAmI(data json.JSON, parameters []string, session *orm.Session) response.Response {
	var currentUser interface{}
	if session.User.ID != 0 {
		currentUser = session.User
	} else {
		currentUser = nil
	}

	return response.New(200, currentUser)
}
