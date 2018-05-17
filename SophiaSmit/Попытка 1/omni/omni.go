package omni

import (
	"bufio"
	"fmt"
	"os"
)

//Функция ввода слова
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func SignIn() {
	var pass string
	var log string
	var login string
	var password string
	var error string

	login = "Stas"
	password = "123"
	error = "Неверный логин или пароль\n"

	log = Scan1()
	pass = Scan1()

	if log == login && pass == password {
		status := true
		dateTimeEnd := 10
		fmt.Print("Все ок\n", status, dateTimeEnd)
	} else {
		status := false
		fmt.Print(error, status)
	}

}
