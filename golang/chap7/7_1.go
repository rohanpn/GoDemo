
	// if we use slice same memory is shared by both the arrays
	// so to allocate memory differently use new()
	
package main
import "fmt"

func main(){

	var arr [5]int
	arr[0]=10
	arr[1]=20
	arr[2]=30
	arr[3]=40
	arr[4]=50
	
	arr1 := arr[0:4]

	arr2 := new([5]int)
	
	for i := 0; i< len(arr); i++{
		arr2[i]= arr[i]
	}
	
	fmt.Println("arr 2 is : ",arr2)
	fmt.Println()		
	
	arr2[3]=500	
	fmt.Print("arr 2 is : ",arr2,"\t");	fmt.Printf("arr 2 address is :  %p",arr2)
	fmt.Println()		

	fmt.Println("arr is : ",arr,"\t");	fmt.Printf("arr address is : %p ",&arr[0])
	fmt.Println()		
	
	fmt.Println("arr 1 is : ",arr1)
//	fmt.Println("arr1 address is : ",&arr1[0])



}