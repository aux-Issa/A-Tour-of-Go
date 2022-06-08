package main
import "fmt"

// 配列は固定長，この場合長さ3の配列となる
var a [3]string
func main(){
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// curly bracesで初期化
	b := [3]int{1,2,3}
	fmt.Println(b)
}