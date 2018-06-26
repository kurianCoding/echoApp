package handlers

import (
	"github.com/kurianCoding/EchoApp/models"
	"github.com/kurianCoding/EchoApp/utils"
)

//TODO: get list of all users
//TODO: get details of all users

func Login(c echo.Context) error {
	user := &models.User{UserName: "", Email: "", Phone: "", Password: "", DBPassword: ""}
	if err := utils.Validate(user); err != nil {
		return error
	}

}
