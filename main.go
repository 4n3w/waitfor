package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func slowFunction() []byte {

	waitFor, err := strconv.Atoi(os.Getenv("SLEEP_INTERVAL_MS"))
	if err != nil || waitFor <= 0 {
		log.Fatal("SLEEP_INTERVAL_MS environment variable not set or not an integer")
	}

	ch := make(chan []byte, 1)

	go func() {
		time.Sleep(time.Duration(waitFor) * time.Millisecond) // Simulate a slow operation
		ch <- []byte("Response after delay")
	}()

	data := <-ch // This will block until the data is available
	return data
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := slowFunction() // Simulate reading from a slow or blocking source
	fmt.Fprintf(w, string(data))
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//
//func main() {
//	waitFor, err := strconv.Atoi(os.Getenv("SLEEP_INTERVAL_MS"))
//	if err != nil || waitFor <= 0 {
//		log.Fatal("SLEEP_INTERVAL_MS environment variable not set or not an integer")
//	}
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		time.Sleep(time.Duration(waitFor) * time.Millisecond)
//
//		// Write response
//		fmt.Fprintf(w, "Response after %d ms delay.", waitFor)
//	})
//
//	// Get the port number from the environment variable
//	port := os.Getenv("PORT")
//	if port == "" {
//		log.Fatal("PORT environment variable not set")
//	}
//
//	fmt.Printf("Server is running on http://localhost:%s\n", port)
//	if err := http.ListenAndServe(":"+port, nil); err != nil {
//		log.Fatal("Server failed:", err)
//	}
//}
