package main

import (
	"example.com/user"
)

func main() {
	appUser, err := user.New("Ahmad", "Reza Adrian", "04-03-2005")
	if err != nil{
		panic(err)
	}

	appUser.OutputUserInfo()
	appUser.SetFirstName("Adrian")
	appUser.SetLastName("Ahmad")
	appUser.OutputUserInfo()

	admin, err := user.NewAdmin("ahmadadrian324@gmail.com", "123")
	if err != nil{
		panic(err)
	}
	admin.GetFirstName()
}