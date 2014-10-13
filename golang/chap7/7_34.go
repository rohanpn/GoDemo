package main
import "fmt"

func main(){

	s := []byte{'p','o','e','m'}
	s2:= s[2:]
	
	fmt.Println("length of s ",len(s))
	fmt.Println("capacity of s ",cap(s))
	s2[1]='z'
	fmt.Println(s)
	
	fmt.Println("length of s ",len(s2))
	fmt.Println("capacity of s ",cap(s2))
	fmt.Println(s2)

}