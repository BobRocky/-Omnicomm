package omni

func SignIn() string {
	var log string
	log = Scan1()
	pass = Scan1()
	var pass string
	var status bool
	login := "Stas"
	password := "123"
	//error := "Неверный логин или пароль\n"
	if log == login && pass == password {
		status = true
		//dateTimeEnd := 10
	} else {
		status = false
	}

	return log
}
