package main
import "fmt"

func main(){
	
	arr := [5]float64{10.10,20.0,30.0,40.0,50.0}
	fmt.Printf("%T",arr)

	var x float64 = sum(arr) 
	fmt.Println()
	fmt.Println(x)
}

func sum(arr [5]float64) (sum float64){
	
	for i:=0;i < len(arr) ; i++{
	
		sum = sum + arr[i]
		
	}
	return sum

}