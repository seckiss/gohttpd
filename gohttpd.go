package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var port string

	if len(os.Args) == 1 {
		port = "80"
	} else if len(os.Args) == 2 {
		iport, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		port = strconv.Itoa(iport)
	}

	go serveHttp(port)
	select {}
}

func serveHttp(port string) {
	log.Printf("Starting http server on port %s\n", port)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, http.FileServer(http.Dir("."))))

}
