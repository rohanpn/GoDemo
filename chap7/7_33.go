package main
import "fmt"

func main(){

	s := make([]byte,5)
	
	fmt.Println("Before slicing")
	fmt.Println("length of s ",len(s))
	fmt.Println("capacity of s ",cap(s))

	s = s[2:4]
	fmt.Println("After slicing")
	fmt.Println("length of s ",len(s))
	fmt.Println("capacity of s ",cap(s))

	var i byte;

	for i =0;i<2;i++{
		s[i]=i*i
	}
	
	fmt.Println("values of i: ")
	
	for i:=0;i<2;i++{
		fmt.Printf("%d\t",s[i])
	}
}