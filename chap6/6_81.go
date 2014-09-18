package main
import "fmt"

func main(){
	a := [...]string{"a", "b", "c", "d"}
	a[10]="s"
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}	
