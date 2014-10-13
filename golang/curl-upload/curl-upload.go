package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	http.Handle("/", http.HandlerFunc(upload))
	http.ListenAndServe(":9090", nil)
}

func log(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req .reqURI  ", req.RequestURI)
	str := strings.Split(req.URL.Path, "/")
	var fname string
	var key int
	str = str[1:]
	for key = range str {
		if key == len(str)-1 {
			break
		}

		fname += str[key] + "/"
	}
	fname += str[key]

	body, err := ioutil.ReadFile(fname)
	if nil != err {
		fmt.Println("read in log   file at : ", err)
		return
	}

	fmt.Fprintf(w, "%+v\n", string(body))
}

func upload(w http.ResponseWriter, req *http.Request) {

	match, _ := regexp.MatchString("\\w{5}/[\\w|\\d]+.\\w+", req.RequestURI)

	if match {
		log(w, req)
		return
	} else {
		elsePart(w, req)
	}
}

func elsePart(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		fmt.Println("returning because of read req.Body error : ", err)
		return
	}

	str := strings.Split(req.URL.Path, "/")
	fname := str[1:]

	file := randString(5)
	err = os.MkdirAll(file, 0750)
	if nil != err {
		fmt.Println("returned error in Mkdir : ", err)
	}

	filename := file + "/" + fname[0]
	dest, err1 := os.Create(filename)
	if nil != err1 {
		fmt.Println("returning because of create1  error : ", err1)
		return
	}

	retWrite, err1 := dest.WriteString(string(body))
	if nil != err1 {
		fmt.Println("retWrite returns : ", retWrite)
		fmt.Println("write 2  error because of  : ", err1)
	}

	dest.Close()

	path := "http://" + req.Host + "/" + filename

	fmt.Fprintf(w, "%+v\n", path)
	fmt.Fprintf(w, "Data in the File is : \n%+v\n", string(body))
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
