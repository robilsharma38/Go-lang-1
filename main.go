package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robilsharma38/mongodb/router"
)

func main() {
	fmt.Println("API")
	r := router.Router()
	fmt.Println("Testing ....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is started .....")
}
