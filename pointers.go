package main
import "fmt"

func main(){
	i,j := 42, 2701
  // iへの参照を渡している
	p := &i
	// pの値を出力する，p自体は16進数のメモリアドレスを示す
	fmt.Println(*p)

	// 参照先の値が変更されたのでiも変更される
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}