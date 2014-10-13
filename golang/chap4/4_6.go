package main
import "fmt"

func main(){


	 var s string = "asassadSDS   fgdf dfbdf d fbdb"
	var s1 string = "asa\u00123ssadSDS   fgdf dfbdf d fbdb"
	var s2 string = "\u0012asa\u0012ssadSDS   fgdf dfbdf d fbdb"

	fmt.Println(len(s))
	fmt.Println(len(s1))
	fmt.Println(len(s2))
}