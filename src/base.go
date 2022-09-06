package main

import "fmt"
var l = fmt.Printf
var ln = fmt.Println

func main() {
   name := "geekli"
   sex := 0
   age := 18 
   const PI = 3.14 
   p := &name

   baby := []string{"迪丽热巴","古力娜扎"}

   pattern := "姓名：%s,性别：%v,年龄：%v\n"

   l(pattern,name,sex,age)
   ln("数组输出：",baby)
   ln("数组长度输出：",len(baby))
   ln("pi =  ",PI)
   ln(p)
   baby = append(baby,"水域梦仙")
   baby[2] = "true love"
}