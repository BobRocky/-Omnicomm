package main

import (
	"fmt"

	"github.com/BobRocky/omnicomm/SophiaSmit/One/omni"
)

var a bool

func main() {

	a = omni.SignIn()
	if a == false {
		error := "Вы ввели не верный логин или пароль\n"
		fmt.Println(a)
		fmt.Print(error)
	} else {
		dateTimeEnd := 10
		fmt.Println(a)
		fmt.Print(dateTimeEnd)
	}
}
