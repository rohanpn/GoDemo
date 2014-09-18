// fibonacci array
package main
import "fmt"
func main(){

	print("fibonacci series \n")
	
	var arr [50]uint
	
	arr[0]=0;
	arr[1]=1;
	
	for i :=0;i<48 ;i++{
		arr[i+2]= arr[i]+arr[i+1];
	
	}
	fmt.Println("fibonacci series : ",arr)
		
	 	

}

