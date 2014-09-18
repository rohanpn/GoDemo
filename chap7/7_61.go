package main
import "fmt"

func main(){

	arr := [...]int{10,20,30,40,50}
	
	for _,ar := range arr{
		ar	 *=	2
	
		fmt.Println("ar : ",ar)
}
	fmt.Println("arr : ",arr)

}