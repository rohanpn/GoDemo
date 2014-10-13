package main
import "fmt"

func main(){

	var i int
	fmt.Scanf("%d",&i)

	switch i{
		case 0 : fmt.Println("zero")
				fmt.Println("zero")	
		case 1 : fmt.Println("one")
		case 2 : fmt.Println("two")
		case 3 : fmt.Println("three")
		default :fmt.Println("greater than three")
	
	}
}