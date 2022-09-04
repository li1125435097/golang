package main

import "fmt"
var l = fmt.Println

type Student struct{
	name string
	gender int
	age int
}

func (this *Student) toString(){
	var sex string
	if(this.gender == 0){
		sex="男"
	}else{
		sex="女"
	}
	l("姓名是：",this.name," 性别是：",sex," 年龄是：",this.age)
}

func (this *Student) clone(that *Student){
	this.name = that.name
	this.gender = that.gender
	this.age = that.age
}


func main(){
	var zs Student
	var ls Student  	// ls := Student{"李四",0,18}
	ls.name = "李四"
	ls.gender = 0
	ls.age = 18

	ls.toString()
	zs = ls		// zs.clone(&ls)
	zs.name = "张三"
	zs.age = 50
	l(zs,ls)
}