package main
import "fmt"

var a string ="global"

func main(){

	a="local"
	print("before m : ",a,"\n")
	m()
	
	fmt.Println("after m :")
	fmt.Println(a)
	
	n()
	
	fmt.Println("after n :")
	fmt.Println(a)

}

func m(){
	a:=100
	print(a)
	print("\n")
}


func n(){
	print(a)
}


