package main
import "fmt"

func main(){

	var x int =0
	label:
		fmt.Println("x : ",x)
		if x!=14{
			x++;goto label
		}
		print("goto completed!!!\n")
		


}