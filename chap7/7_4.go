package main
import "fmt"

func main(){
//	var arr [2]int
	x,y:= fibo()

	z  := x[y[0]:y[1]]
	fmt.Println("sliced array : ",z)
}

func fibo()  ( arr [10]int ,arr1 [2]int){
	
//		var arr1[2]int
	arr1[0]=7
	arr1[1]=10
//ariarr [50]int
	arr[1]=1
	for i :=0;i<8 ;i++{
		arr[i+2]= arr[i]+arr[i+1];
	
	}
	fmt.Println("arr is : ",arr)
	return 

}