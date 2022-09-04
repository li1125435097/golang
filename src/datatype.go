package main

import "fmt"
var l = fmt.Println
const (
	Monday = iota
	Tuesday = iota
	Wednesday = iota
	Thursday = iota
	Friday = iota
)


func main()  {
	defer l("rabbit")
	var any interface{}
	var dict map[string]string
	dict = make(map[string]string,1)

	l(Monday)
	any = "sdf";l(any)
	any = []int{12,33};l(any)
	dict["男"]="man";l(dict)
	dict["女"]="woman";l(dict)
	fmt.Printf("%T\n",any)
}