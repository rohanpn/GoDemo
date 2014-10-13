package main
import "fmt"

func main(){
	s:= []byte{1,2,3,4,5,6,7,8,9,10}
	
	s1,s2 :=s[1:4],s[3:8]
	
	fmt.Println("s : ",s)
	s = s[0:4]

	fmt.Println("s : ",s)

	fmt.Println("s1 : ",s1)
	fmt.Println("s2 : ",s2)
	
	s1 = s1[0:len(s1)]
	
	fmt.Println("s1 : ",s1)
	fmt.Println("s2 : ",s2)

	s = s[0:7]

	fmt.Println("s : ",s)

}