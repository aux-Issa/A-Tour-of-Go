package main
import "fmt"

func main(){
	names := [4]string{
		"Aoki",
		"Ishida",
		"Ueda",
		"Endo"
	}
	first_two = names[0:2]
	last_two = names[2:]
	fmt.Println(first_two, last_two)

	first_two[1] = "XXX"
	last_two[1] = "XXX"
	fmt.Plintln(first_two, last_two, names)
}