package main

import (
	"log"

	"github.com/murasame29/distributed_services_with_go_memo/letsgo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
