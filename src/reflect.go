package main

import(
	"fmt"
	"reflect"
)
var l = fmt.Println
var types = reflect.TypeOf
var vals = reflect.ValueOf

type Man struct{
	name string
	meat bool
}
func (this Man) greeting() {
	l(this.name,": hello, adults and teenagers")
}

func main()  {
	zs := Man{"张三",true}
	zs.greeting()
	l(zs)
	l(types(zs))
	l(vals(zs))
}