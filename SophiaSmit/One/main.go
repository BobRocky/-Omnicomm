package main

import (
	"fmt"

	"github.com/BobRocky/omnicomm/SophiaSmit/One/omni"
)

var a bool

func main() {

	a = omni.SignIn()
	fmt.Println(a)
}
