package main

import "fmt"
var l = fmt.Println

type Animal interface{
	sayHi() 
}

type Cat struct{
	name string
}

func (this *Cat) sayHi(){
	l(this.name,"miao miao~ ")
}

type Dog struct{
	name string
}

func (this *Dog) sayHi(){
	l(this.name,"wang wang~ ")
}

func main()  {
	var any Animal
	any = &Cat{"小花"}
	any.sayHi()

	any = &Dog{"小黄"}
	any.sayHi()
}