package main

import "fmt"

func main(){
	sum_while := 1
	sum_for := 1
	for sum_while < 100 {
		sum_while += sum_while
	}
	for i:=1 ; i < 100 ; i++{
		sum_for += i
	}
	fmt.Println(sum_while,sum_for)
}