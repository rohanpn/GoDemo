package main
import "fmt"

func main(){
	
	s := ""
	for  ; s!="aaaaa";{
		fmt.Println("value of s : ",s)
		s = s +"a"
	}

	a:=1
	b:=9
	b+=a
	print("a is %v *** b is *v ",a,b)

}