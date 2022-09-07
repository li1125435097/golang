package main

import "fmt"
var l = fmt.Println

func sum(a int,b int,c chan int) {
	c <- a+b  
}

func fibonacci(n int, c chan int){
	s := 1
	e := 1
	temp := 0

	for i := 0; i < n; i++ {
		temp = e
		e = s+e
		s = temp
		c <- e
	}
	close(c)
}

func main()  {
	ch := make(chan int)
	ch2 := make(chan int,10)
	
	

	l("--------------")
	go fibonacci(cap(ch2),ch2)
	go sum(1,3,ch)
	for i := range ch2{ l(i) }
	x := <- ch;l(x,"first")
}