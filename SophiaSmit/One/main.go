package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/BobRocky/omnicomm/SophiaSmit/One/omni"
)

var log string
var pass string
var a bool

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {

	log = Scan1()
	pass = Scan1()
	a = omni.SignIn()
	fmt.Println(a)
}
