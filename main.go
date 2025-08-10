package main

import (
	"fmt"
	"net/http"
	"io"

	// "database\sql"
	"log"
	// _ "github.com/go-sql-driver/mysql"
	// "errors"
)

const urlHui = "https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities.json"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"check ping"}`)
	})

	mux.HandleFunc("/moex", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		resp, err := http.Get(urlHui)
		if err != nil {
			fmt.Println(err)
		}
		
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")

		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Printf("write to client error: %v", err)
			return
		}
	})

	log.Println("listening")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}