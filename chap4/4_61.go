package main
import (
		"fmt"
		"strings"
		)
func main(){

	var str string = "this is the example of the string prefix"

// HasPrefix() checks whether the string has a part of  is as a start of the string
	fmt.Println(strings.HasPrefix(str,"is"))
	
	// Contains()
	fmt.Println(strings.Contains(str,"the"))
	
}