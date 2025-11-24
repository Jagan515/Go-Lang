package main

import (
	"fmt"

	"github.com/Jagan515/Go-Lang/auth"
	"github.com/Jagan515/Go-Lang/user"
	"github.com/fatih/color"
)

func main() {
	auth.Loginwuthcred("jagan", "password123")
	auth.Signupwithcred("jagan", "password123", "jagan@example.com")

	token := auth.GetSessionToken("user123")
	println("Session Token:", token)

	user := user.User{
		Username: "jagan",
		Password: "password123",
		Email:    "jagan@example.com",
	}
	fmt.Println("User Email:  ", user.Email)        // use " go get repo link" to get the package from github
	color.Cyan("User Username:  %s", user.Username) // use " go mod tidy" to add the dependency to go.mod file autmomatically
	color.Green("User Password: %s", user.Password)

}
