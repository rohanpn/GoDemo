package main
import (
		"fmt"
		"strings"
		)
		
func main(){
	str :="Hi, I am rohan Hi"
	str1 :="bye"
	fmt.Println("i : ",strings.Index(str,"I"))

	fmt.Println("first hi : ",strings.Index(str,"Hi"))
	fmt.Println("last hi : ",strings.LastIndex(str,"Hi"))
	fmt.Println("burger : ",strings.Index(str,"n"))

	strings.Replace(str,str,str1,len(str1))
	fmt.Println(str)

}