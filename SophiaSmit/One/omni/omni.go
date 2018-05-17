package omni

import (
	"bufio"
	"fmt"
	"os"
)

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func SignIn() bool {
	var log string
	var pass string
	var status bool
	log = Scan1()
	pass = Scan1()

	login := "Stas"
	password := "123"
	//error := "Неверный логин или пароль\n"
	if log == login && pass == password {
		status = true
		//dateTimeEnd := 10
	} else {
		status = false
	}

	return status
}
