package main
import "fmt"

func main(){
	
	g := func() {
			fmt.Println("hello...")
			return 
		}
	g()
	print("world\n")

	const n =true
	fmt.Printf("tpe of n is  %T\n",n)
	fmt.Println(f())

}
func f() (ret int) {
	
	defer func() {
		ret++
		fmt.Println(ret)
	}()
	
	return 1
}
