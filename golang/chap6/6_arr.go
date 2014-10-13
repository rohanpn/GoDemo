package main
import "fmt"

func main(){

	var arr1 [5]int
	var arrn = new([5]int)
	arrn[1] =20
	
	
	fmt.Printf("arr1 : %d\n",arr1)
	fmt.Printf("arrn : %p \n",arrn)
	fmt.Println("arrn : \n",arrn[0],arrn[1],arrn[2],arrn[3],arrn[4])	
	arr3 := arrn[3]

	fmt.Printf("arr3 : %d\n",arr3)
}