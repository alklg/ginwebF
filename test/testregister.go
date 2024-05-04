package test

import (
	"ginweb/internal/db"
	"ginweb/internal/models"
)

func TestRegister(user models.User) {

	db.InsertUserIntoDB(&user)

}
