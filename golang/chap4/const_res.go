package main
import (
		"fmt"
	)		
func main(){

	type ByteSize float64
	const(
		_ = itoa
		KB  ByteSize = 1<<(10*itoa)
		MB
		GB
	)	
	
	fmt.Println(KB,MB,GB,ByteSize)

}