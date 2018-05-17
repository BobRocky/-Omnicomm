package main

import (
	"bufio"
	"fmt"
	"github.com/BobRocky/omnicomm/SophiaSmit/One/omni"
)

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
func main() {
	var log string
	var pass string

	log = Scan1()
	pass = Scan1()
	proverka := SignIn()
	fmt.Print(proverka)

}