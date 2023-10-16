package server

import (
	"log"

	"fmt"
	"net/http"
)

func StartServer(port string) {
	fmt.Println("Starting application at port : ", port)

	handler, err := New()
	if err != nil {
		log.Fatal("Failed to init server. Error : ", err)
	}
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
