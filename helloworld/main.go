// A sample web service in Go to test the configuration of
// the Linux box, Apache proxying a web site, and configuring
// a system.d service to restart on boot. This web service
// also tests reading an ENV variable and a file from disk.
//
// Build with:
//		GOOS=linux go build -o helloworld_linux
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := flag.Int("port", 9001, "Listening port")
	flag.Parse()

	// Start the web server
	webAddress := fmt.Sprintf("localhost:%d", *port)
	log.Printf("Listening for requests at: http://%s", webAddress)
	log.Printf("Environment DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("Endpoints:")
	log.Printf("\t /hello/")
	log.Printf("\t /file/")
	log.Printf("\t /panic/")

	http.HandleFunc("/file/", filePages)
	http.HandleFunc("/hello/", helloPages)
	http.HandleFunc("/panic/", panicPages)
	http.HandleFunc("/", rootPages)
	err := http.ListenAndServe(webAddress, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}

func rootPages(resp http.ResponseWriter, req *http.Request) {
	log.Printf("root endpoint")
	msg := fmt.Sprintf("root for DB_USER: %s", os.Getenv("DB_USER"))
	fmt.Fprint(resp, msg)
}

func helloPages(resp http.ResponseWriter, req *http.Request) {
	log.Printf("hello endpoint")
	msg := fmt.Sprintf("hello at %s", utcNow())
	fmt.Fprint(resp, msg)
}

func filePages(resp http.ResponseWriter, req *http.Request) {
	log.Printf("file endpoint")
	filename := "./hello.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(resp, string(bytes))
	} else {
		fmt.Fprint(resp, err)
	}
}

func panicPages(resp http.ResponseWriter, req *http.Request) {
	log.Printf("panic endpoint")
	panic("we hit the panic endpoint")
}

func utcNow() string {
	t := time.Now().UTC()
	s := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return s
}
