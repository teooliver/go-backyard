package simpleserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/good", func(http.ResponseWriter, *http.Request) {
		log.Println("Good!")
	})

	http.ListenAndServe(":9090", nil)
}
