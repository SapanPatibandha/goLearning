package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"} // this should be pointer
var mut sync.Mutex             // this should be pointer.

func main() {
	fmt.Println("Welcome to goRutin")

	// go greeter("hellow")
	// greeter("World")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	//
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("error during call")
	} else {

		// this is something like lock object in C#
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code from website %s \n", res.StatusCode, endpoint)
	}

}
