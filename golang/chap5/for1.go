package main
import "fmt"

func main(){

	i:=1
	for i<=10{
		fmt.Print(" ",i)
		if  i % 2 ==  0{
			fmt.Println("div by 2")
		}else if i%3 ==0 {
			fmt.Println("div by 3")
	
		}
		i+=1
	}
}
