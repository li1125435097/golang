package main
import "fmt"
var l = fmt.Println


func main()  {
	arr := []int{4,5,1,2,3}
	dict := make(map[string]int)
	dict["马里奥"]=1
	dict["佩奇"]=2
	dict["喜羊羊"]=3
	len:=len(arr)

	// 普通循环
	for i := 0; i < len; i++ { l(arr[i])	}
	l("-----------------")

	// range迭代
	for i,v := range arr { l(i,v) }
	for i,v := range dict { l(i,v) }

	l(dict.len)
}