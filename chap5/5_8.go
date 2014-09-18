package main
import "fmt"

var pos int
func main(){

	
//	var arr[20][10] byte
	
	for i:=0;i<5;i++{
		for j:=0;j<10;j++{
			if j == 0 || j ==9 ||  i==0 || i ==4{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}	
		};
		fmt.Printf("\n")
	}

	fmt.Print(pos)
	str := "string"
//	sptr := &str
	for pos, val:=range str{
		
		fmt.Printf("%c\t",val)
//		val = pos
	}
}