package main

import "fmt"
var l = fmt.Println

func main()  {
	dict := make(map[string]string)
	dict["山西"] = "陈醋"
	dict["广东"] = "佛跳墙"
	dict["北京"] = "炸酱面"

	l(len(dict))
	data,ok := dict["山西"];l(data,ok)
	data1,ok1 := dict["河南"];l(data1,ok1)

	delete(dict,"山西");l(dict)
}