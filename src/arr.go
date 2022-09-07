package main

import "fmt"
var l = fmt.Println

func main()  {
	// 声明
	var(
		arr1 [1]int
		arr2 [2]int
		sc []int
		sc1 = make([]int, 3)
		sc2 = []int{1,54,5,4}
		sc3 = []int{4:1,3:2}
	) 
	

	arr1[0] = 1;l(arr1)
	arr2[0] = 1;l(arr2)
	arr2[1] = 2;l(arr2)
	sc = append(sc,1);l(sc,len(sc))
	sc1[1] = 1;l(sc1,len(sc1),cap(sc1))
	sc1[2] = 2;l(sc1,len(sc1),cap(sc1))
	l(sc2,len(sc2),cap(sc2))
	l(sc3,len(sc3),cap(sc3))

}