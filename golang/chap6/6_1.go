package main
import "fmt"

func main(){

	a:=10
	b:=20
	
	var sum,diff int
	
	sum,diff = arithematic(a,b)

	fmt.Println(sum,diff)
	
}
func arithematic(a,b int)  (add,sub int){

	add = a+b; sub = a-b
	return 
}