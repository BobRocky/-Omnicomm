package main

import "fmt"
import x "github.com/BobRocky/omnicomm/SophiaSmit/Test/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := x.Average(xs)
	fmt.Println(avg)
}
