 package main

// import (
// 	"fmt"
// 	"html"
// 	"log"
// 	"net/http"
// )

//  BASIC SERVER

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// 	})

// 	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hi")
// 	})

// 	log.Fatal(http.ListenAndServe(":8081", nil))

// }

// //(2) SECOND WAY
// func helloPrint(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

// }
// func main() {

// 	http.HandleFunc("/", helloPrint)

// 	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hi")
// 	})

// 	log.Fatal(http.ListenAndServe(":8081", nil))

// }

//(3)SERVER WITH MUTEX
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
