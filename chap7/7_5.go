package main
import "fmt"

func main(){
	s := []byte {'a','b','c','d','e','f','g','h'}
	
	s1 := s[2:5]
	
	slice := Append(s1,s)

		fmt.Println(slice)
}

func Append(slice ,data []byte) []byte{

	l:= len(slice)
	fmt.Println("data : ",data)
	fmt.Printf("length of slice : %d\n",l)
	fmt.Printf("Capacity of slice : %d\n",cap(slice))
	fmt.Printf("length of  data: %d\n",len(data))
	fmt.Println("new length for slice",(l+len(data))*2)
	
	if l + len(data) > cap(slice){
		newSlice:=make([]byte,(l+len(data))*2)
		copy(newSlice,slice)
		slice = newSlice


		fmt.Println(slice)
	}
	
	slice = slice[0:((l+len(data))*2)]
	for i,c := range data{
		slice[l+i]=c
	}
	return slice
}