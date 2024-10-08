package main

import (
	"fmt"
	handler "golangnext/api"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/chatnew", enableCORS(handler.ChatNew))
	http.HandleFunc("/chatlist", enableCORS(handler.ChatList))
	http.HandleFunc("/chatsend", enableCORS(handler.ChatSend))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
