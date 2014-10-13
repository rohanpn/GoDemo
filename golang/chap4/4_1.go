package main
import "fmt"

var a int =10

func main(){

	m()
	n()
	o()
	fmt.Println(a)
}

func m(){ 
	print(a)
}
func n(){
	a:="inside n"
	fmt.Println(a)
}
func o(){
	a=20
	fmt.Println("o : ",a)
}