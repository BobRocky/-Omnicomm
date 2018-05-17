package omni

import "fmt"

func SignIn(string) {
	var login string
	var password string
	var error string

	login = "Stas"
	password = "123"
	error = "Неверный логин или пароль\n"
	if log == login && pass == password {
		status := true
		dateTimeEnd := 10
		fmt.Print("Все ок\n", status, dateTimeEnd)
		return status, dataTimeEnd
	} else {
		status := false
		fmt.Print(error, status)
		return status, error
	}

}
