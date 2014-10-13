package main
import "fmt"

func main(){
	var by [7]byte 
	by[0]=1
	by[1]=12
	by[2]=222
	by[3]=32
	by[4]=43
	by[5]=54
	by[6]=65
	
	
	bt2 := by[2:5]
	bt3 := by[5:7]
	bt4 := by[3:]

	fmt.Println("bt2 : ",bt2)
	fmt.Println("bt3 : ",bt3)
	fmt.Println("bt4 : ",bt4)

}