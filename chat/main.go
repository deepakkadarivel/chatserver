package main

import (
	"log"
	"net/http"
	"flag"
	"github.com/gorilla/mux"
	"encoding/json"
	"chatserver/chat/auth"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	var addr = flag.String("addr", ":8000", "The addr of the application.")
	flag.Parse()
	router := mux.NewRouter()
	r := newRoom()
	router.Handle("/", r)
	go r.run()
	// This route is always accessible
	router.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello from a public endpoint! You don't need to be authenticated to see this.",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))

	// This route is only accessible if the user has a valid access_token with the read:messages scope
	// We are wrapping the checkJwt middleware around the handler function which will check for a
	// valid token and scope.
	router.Handle("/api/private", auth.CheckJwt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello from a private endpoint! You need to be authenticated and have a scope of read:messages to see this.",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	})))
	log.Println("Started server and listining at port : ", *addr)
	if err := http.ListenAndServe(*addr, router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
