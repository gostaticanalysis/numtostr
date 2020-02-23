package a

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	s := fmt.Sprint(rand.Int())  // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = strconv.Itoa(rand.Int()) // OK
	fmt.Println(s)
	s = fmt.Sprint(true)   // OK
	s = fmt.Sprint(10, 11) // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`

	v := 10
	s = fmt.Sprint(v) // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
}
