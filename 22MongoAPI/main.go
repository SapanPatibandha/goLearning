package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SapanPatibandha/goLearning/22MongoAPI/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":6000", r))
	fmt.Println("Listening at port 6000 ...")
}
