package main

import (
    "fmt"
    "net/http"
    "os"
//    "log"
    "io/ioutil"
	"strings"
)

type handle struct {
    count int
}

func (o *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    //  fmt.Println(o.count)
    fmt.Fprintf(w, "hello %d", o.count)
    o.count++
}

func main() {
    h := handle{}

	args := os.Args[1]
	file_ptr,err  := os.Open(args)
	if nil != err{
		fmt.Println("error opening file")
	}
	
	fmt.Println("file+ptr is :",file_ptr)

	
	fData, err := ioutil.ReadFile(args)
    fmt.Println("fData : ",string(fData))
	
	
	arr := strings.SplitAfterN(string(fData),"=",4)
//	fmt.Println(strings.Trim(strings(fData),"\""))
	arr1 := arr[1][:len(arr[1])-2]
	fmt.Println("a0: ",arr[0])
	fmt.Println("a1: ",arr[1])
	fmt.Println("arr1 :",arr1)
   	http.Handle("/count", &h)

    fmt.Println("starting server...")
   	http.ListenAndServe(string(arr1), nil)
 

}
