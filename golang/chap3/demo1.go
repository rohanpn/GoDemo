package main
import "fmt"
func main(){
	
	var c int
	var f int
	
	fmt.Println("Enter farenheit value :  ")
	fmt.Scanf("%d",&f)
	
	c = (f-32) * 5/9

	fmt.Println("value before converting : ",f)
	fmt.Println("value after converting : ",c)
}