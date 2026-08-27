package main

import (
	"log"

	"github.com/Shreyas9468/proglog/internal/server"
)

func main() {
	httpSrv := server.NewHTTPServer(":8080")
	log.Fatal(httpSrv.ListenAndServe())
}
