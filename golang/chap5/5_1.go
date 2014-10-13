package main
import "fmt"

func main() {
	
	for i:=0; i<10; i++ {
		fmt.Printf("%p\n", &i)
	}
	
	
	str2 := "???"
	
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	 for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

}
