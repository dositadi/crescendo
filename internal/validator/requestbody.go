package validator

import "regexp"

func ValidFormValues(username, email, password string) (bool, string) {
	switch {
	case username == "":
		return false, "Username field cannot be empty"
	case email == "":
		return false, "Email field cannot be empty"
	case password == "":
		return false, "Password field cannot be empty"
	case email != "":
		if !validateEmail(email) {
			return false, "Invalid email address format. Example: example@gmail.com"
		}
	case password != "":
		if len(password) < 8 {
			return false, "Password should be at least 8 characters long"
		}
	}
	return true, ""
}

func validateEmail(email string) bool {
	valid, _ := regexp.Match(`[a-zA-Z]+[a-zA-Z0-9-_.]*@[a-z]{8}[.]{1}[a-z]{3}`, []byte(email))

	return valid
}
