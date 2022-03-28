// A sample web service in Go to test the configuration of
// the Linux box, Apache proxying a web site, and configuring
// a system.d service to restart on boot. This web service
// also tests reading an ENV variable and a file from disk.
//
// Build with:
//		GOOS=linux go build -o helloworld_linux main.go
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
	log.Printf("Environment DB_USER: %s", envVariable())
	log.Printf("Endpoints:")
	log.Printf("\t /hello/")
	log.Printf("\t /file/")
	log.Printf("\t /panic/")

	http.HandleFunc("/file/", filePage)
	http.HandleFunc("/hello/", helloPage)
	http.HandleFunc("/panic/", panicPage)
	http.HandleFunc("/", rootPage)
	err := http.ListenAndServe(webAddress, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}

func rootPage(resp http.ResponseWriter, req *http.Request) {
	log.Printf("root endpoint")
	msg := fmt.Sprintf("root page, DB_USER: %s\r\n", envVariable())
	fmt.Fprint(resp, msg)
}

func helloPage(resp http.ResponseWriter, req *http.Request) {
	log.Printf("hello endpoint")
	const time_format_now string = "2006-01-02 15:04:05.000" // yyyy-mm-dd hh:mm:ss.xxx
	now := time.Now().Format(time_format_now)
	msg := fmt.Sprintf("hello at %s\r\n", now)
	fmt.Fprint(resp, msg)
}

func filePage(resp http.ResponseWriter, req *http.Request) {
	log.Printf("file endpoint")
	filename := "./hello.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Fprint(resp, string(bytes))
	} else {
		fmt.Fprint(resp, err)
	}
}

func panicPage(resp http.ResponseWriter, req *http.Request) {
	log.Printf("panic endpoint")
	panic("we hit the panic endpoint")
}

// Look for a test ENV variable
func envVariable() string {
	value := os.Getenv("DB_USER")
	if value == "" {
		return "(not set)"
	}
	return value
}
