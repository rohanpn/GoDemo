package main

import (	
		"fmt"
		"io/ioutil"
		"math/rand"
		"net/http"
		"os"
		"regexp"
		"strings"
		"log"
	)

var letters = []rune("abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":9090", nil)
}

func getFileName(url string) string{

	str := strings.Split(url, "/")
    str = str[1:]
    fname := strings.Join(str ,"/")

	return fname
}

func sendFile(w http.ResponseWriter, req *http.Request) {
	
	fname := getFileName(req.URL.Path)
	_, err := ioutil.ReadFile(fname)
	if nil != err {
		log.Fatal(err)
		return
	}

}

func handler(w http.ResponseWriter, req *http.Request) {

	match, _ := regexp.MatchString("\\w+/[\\w|\\d]+.\\w+", req.RequestURI)

	if match {
		sendFile(w, req)
		return
	} else {
		uploadFile(w, req)
	}
}

func sendLink(w http.ResponseWriter, host string,filename string) bool{

	path := "http://"+host+"/"+filename
	fmt.Fprintf(w,"%+v\n",path)
	return true

}
func uploadFile(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if nil != err {
		log.Fatal(err)	
 	return
	}

	str := strings.Split(req.URL.Path, "/")
	fname := str[1:]

	dir := randString(5)
	err = os.MkdirAll(dir, 0750)
	if nil != err {
		log.Fatal(err)
	}
	
	filename := dir + "/" + fname[0]
	
	err = ioutil.WriteFile(filename,body,0750)
	if nil != err {
		log.Fatal(err)
	}
	sendLink(w,req.Host,filename)

}


func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
