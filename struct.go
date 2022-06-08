package main
import "fmt"


type Vertex struct {
	X int
	Y int
}

var (
		var1 = Vertex{1,2}
		var2 = Vertex{X:1}
		var3 = Vertex{}
		p = &Vertex{1, 2}
)
func main(){
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	fmt.Println(var1, var2, var3, p)
}