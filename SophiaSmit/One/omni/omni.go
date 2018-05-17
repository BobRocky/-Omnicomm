package omni

import "fmt"

func SignIn() {
	var login string
	var password string
	var error string
	var log string
	var pass string

	login = "Stas"
	password = "123"
	error = "Неверный логин или пароль\n"
	if log == login && pass == password {
		status := true
		dateTimeEnd := 10
		fmt.Print("Все ок\n", status, dateTimeEnd)
		return string(status, dataTimeEnd)
	} else {
		status := false
		fmt.Print(error, status)
		return string(status, error)
	}

}
