// https://go-tour-jp.appspot.com/moretypes/7
package main

import "fmt"

func main(){
	primes := [4]int{1,2,3,4}
	// 配列は固定長であるのに対して
	// スライスは可変長である
	var slice []int = primes[1:4]
	fmt.Println(slice)
}