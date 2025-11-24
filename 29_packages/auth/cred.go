package auth

import (
	"fmt"
)

// Captical letter means public functionglobal access
// small letter means private function limited to package only and folder
func getusername(name string) string {
	return name + "W_auth" // not accesssible outside this package
}

func Loginwuthcred(username string, password string) bool {
	// logic to check username and password against database
	fmt.Println("logged in successfully ", getusername(username))
	return true
}
func Signupwithcred(username string, password string, email string) bool {
	// logic to store username , password and email to database
	fmt.Println("signed up successfully ", getusername(username))
	return true
}
