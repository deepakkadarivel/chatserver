package main

import (
	"log"
	"net/http"
	"flag"
	"github.com/gorilla/mux"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()
	router := mux.NewRouter()
	r := newRoom()
	router.Handle("/", r)
	go r.run()
	log.Println("Started server and listining at port : ", *addr)
	if err := http.ListenAndServe(*addr, router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`
	//		<html>
	//			<head>
	//			  <title>Chat</title>
	//			</head>
	//			<body>
	//			  Let's chat!
	//			</body>
     // 		</html>
	//	`))
	//})
	//http.Handle("/", r)
	//
	//r.run()
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal("ListenAndServe:", err)
	//}
}
