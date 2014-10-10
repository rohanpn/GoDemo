package main


import(
		"fmt"
		"io/ioutil"
		"net/http"
		"strings"
		"os"
//		"path/filepath"
//		"runtime"
//		"path"
		"math/rand"
//		"io"
	)

	func init(){
		os.Mkdir("file",0750)
	}
func main() {
    	
	 http.Handle("/",http.HandlerFunc(upload))
	http.Handle("/file/",http.HandlerFunc(log))  	 
	http.ListenAndServe(":9090", nil)
}


func log(w http.ResponseWriter,req * http.Request){
	
	fmt.Fprintf(w,"/file/ logged..!!!")
	str := strings.Split(req.URL.Path,"/")
    fname := str[1]+"/"+str[2]

//	fmt.Println("----------------------")
 //  fmt.Printf("filename in /file : %+v\n ",fname)

	 body,err := ioutil.ReadFile(fname)
      if nil != err{
            fmt.Println("read   file at : ",err)
            return
      }
      //  fmt.Fprintf(w," path given is : %+v\n",path)
        fmt.Fprintf(w," read returned : %+v\n",string(body))

}



func upload(w http.ResponseWriter,req * http.Request){

	 	reqInfo(req)

	fmt.Println("host name is : ",req.Host)
		
	body ,err := ioutil.ReadAll(req.Body)
		if nil != err{

			fmt.Println("returning because of read req.Body error : ",err)
			
			return
		}

		
	fmt.Fprintf(w,"printing the data(req.Body) %+v\n",string(body))

	
	////////////////// get the file name frm the url ////////////////

	str := strings.Split(req.URL.Path,"/")
	fname := str[1:]

	
	fmt.Println("filename is : ",fname[0])
	
	

	////////////////// create a file for that name///////////	
	
	filename := "file/"+fname[0]
	fmt.Println("----------------------------")
	fmt.Println("file path created is  :  ",filename)
	dest ,err1 := os.Create(filename)
	if  nil != err1{

          fmt.Println("returning because of create1  error : ",err1)

          return
	}
	
	
	
	/////////////////// write to the file created on server	///////////////////

	fmt.Println("before copying : ",string(body))
	retWrite,err1 := dest.WriteString(string(body))
		if nil != err1{
			fmt.Println("write 2  error because of  : ",err1)
		}
	
		fmt.Printf("contents copied to file... : %+v\n",string(body))
		
		fmt.Println(retWrite)	
		dest.Close()
	

	//////////////////// url of the file ///////////////

	path := "http://"+req.Host+"/"+filename
		
	fmt.Fprintf(w,"%+v\n",path)
//	fmt.Println("path is : ",path)
/*
	//////////////////////// open the file //////////////////

	
	fp,err := os.Open(filename[0])
	if nil != err{
		fmt.Println("error opening at : ",err)
	}

			
	/////////// read data from the file to display on browser /////////////

	/*
		body,err1 = ioutil.ReadFile(filename[0]) 
		if nil != err1{
			fmt.Println("read   file at : ",err)
			return
		}
	  	fmt.Fprintf(w," path given is : %+v\n",path)
		fmt.Fprintf(w," read returned : %+v\n",string(body))

/*
	//	var  bytes []byte
		n,er :=  fp.Read(bytes)	
		fmt.Println("err0r is : ",er)
		fmt.Println("read from files : ",n)
*/
		bytes,_ := ioutil.ReadAll(req.Body)
	 
		fmt.Fprintf(w,"data is : ",string(bytes))
//		fp.Close()
//		fmt.Println("random name is : ",randSeq(5))

	 
}

func reqInfo(req *http.Request){

	 fmt.Println("method :",req.Method)

		fmt.Println("\n URL details..\n")
//	fmt.Println("URL scheme is : ",req.URL.Scheme)
//	fmt.Println("URL Opaque is : ",req.URL.Opaque)
/*	
**/     // fmt.Println("URL is :",*(req.URL))
		fmt.Println("URL host is : ",string(req.URL.Host))
		fmt.Println("URL path is : ",req.URL.Path)
	fmt.Println("\n")
  
	
	   fmt.Println("FileSize : ", req.FormValue("filesize"))
    
	  fmt.Println("Header : ",req.Header)
      fmt.Printf("Body : %v\n",req.Body)
      fmt.Println("Encoding : ",req.TransferEncoding)


      fmt.Println("Post form : ",req.PostForm)
      fmt.Println("RemoteAddr : ",req.RemoteAddr)
      fmt.Println("RequestURI : ",req.RequestURI)




}
func visit(path string, f os.FileInfo, err error) error {
 	 fmt.Printf("Visited: %s\n", path)
 	 return nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

