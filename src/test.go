package main

import "fmt"
import "reflect"
var l = fmt.Println

func main()  {
	str := "alsdfi"
	num := 123
	arr := [5]int{4554,45,54}

	var cmd interface{}
	cmd = "123"
	v,ok := cmd.(string)
	l(v,ok)
	fmt.Printf("%T\n",reflect.TypeOf(str).String())
	l(reflect.TypeOf(num))
	l(reflect.TypeOf(arr))
}