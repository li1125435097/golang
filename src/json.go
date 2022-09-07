package main

import "fmt"
import "encoding/json"
var l = fmt.Println

func main()  {
	arr := []int{54,456,456,56,45}
	fill := make([]int,4)
	buf,_ := json.Marshal(arr)
	err := json.Unmarshal(buf,&fill)
	l(string(buf))
	l(err,fill)
}