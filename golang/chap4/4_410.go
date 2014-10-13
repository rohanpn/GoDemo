package main
import (
		"fmt"
		"math/rand"
		"time"
		)
		
func main(){
	
	for i := 0; i<10; i++{
	
		a:= rand.Intn(200)
		fmt.Println("a : ",a)
	}
		
		fmt.Println("Current Time : \t",time.Now())


	var ch byte = '\x41'
		fmt.Printf("ch : %c \n",ch)
}