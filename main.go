package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var contentType *string
var body *string

func main() {
	X := flag.String("X", "", "HTTP method.")
	body = flag.String("D", "", "Content that you want to post.")
	contentType = flag.String("H", "application/json", "Content type.")
	flag.Parse()
	var uri string = os.Args[len(os.Args)-1]
	if os.Args[1] == "help" {
		fmt.Println("Welcome to the Gorl guide.")
		fmt.Println("Usage: gorl [options] <url>")
		fmt.Println("Help: gorl help")
		fmt.Println("--help: Help about command.")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("+---------------------------+")
		fmt.Println("GET: gorl -X GET <url>")
		fmt.Println("POST: gorl -X POST -D <content> <url>")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("+---------------------------+")
		fmt.Println("-X: HTTP method.")
		fmt.Println("-D: Content that you want to post.")
		fmt.Println("+---------------------------+")
		fmt.Println("Project by Haume")
		return
	}

	switch strings.ToUpper(*X) {
	default:
		fmt.Println("try 'gorl --help' for more information.")
		return
	case "GET":
		Get(uri)
		return
	case "POST":
		Post(uri, *body)
		return
	}
}

func Get(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("FETCH ERROR")
		return
	}
	//defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("BODY READ ERROR")
		return
	}
	fmt.Println(string(body))
}
func Post(url string, body string) {
	res, err := http.Post(url, *contentType, strings.NewReader(body))
	if err != nil {
		fmt.Println("FETCH ERROR")
		return
	}
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("BODY READ ERROR")
		return
	}
	fmt.Println(string(resbody))
}
