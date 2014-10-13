package main
import "fmt"

func main(){

	for i:=0;i<50;i++{
	
		switch{
		
			case i%5 ==0  && i%3 ==0 : fmt.Println();fmt.Println("fizzbuzz")

			case i%3 == 0: fmt.Println();fmt.Println("fizz")
		
			case i%5 ==0 : fmt.Println();fmt.Println("buzz")
		
		
			default : fmt.Printf("%d \t",i)
		}
	}

}