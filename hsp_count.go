package main

import (
	"strings"
	"fmt"
	"net/http"
	"os"
	"io"
	"bufio"
	"log"
	"strconv"
)

type handle struct {
	count int
}

func (o *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	
	fmt.Fprintf(w, "hello %d\n", o.count)
	o.count++
}


type Options map[string]string

func scanFile(f io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func linesToOpts(lines []string) (Options,int) {
	opts := make(Options)
	count :=0
	var k string 	
	for _, l := range(lines) {
			
		fmt.Println("l lto : ",l)
		if count ==1{
			k =l
			break;
		} 
		parts := strings.Split(l, "=")
		opts[parts[0]] = parts[1]
		count++	
	}

	i,err := strconv.Atoi(k)
	return opts,i	
} 
func getOpts() (Options,int) {
	if len(os.Args) < 2 {
		log.Fatal("First argument should be config file")
	}
	cf, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer cf.Close()
	lines, err := scanFile(cf)
	optsmap,i := linesToOpts(lines)	
	return optsmap,i		
}

func startServer(port string,Count int) {
	h := handle{count: Count }
	http.Handle("/count", &h)
	port_string := ":" + port

	log.Println("Starting server on %s", port_string)
	http.ListenAndServe(port_string, nil)
}

func main() {
	options,i := getOpts()
	port, ok := options["port"]
	fmt.Println("i : ",i)	
	if !ok {
		log.Fatal("Port not found in config file")
	}
	startServer(port,i)
}


