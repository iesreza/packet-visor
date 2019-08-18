package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iesreza/packet-visor/pvisor"
	"io/ioutil"
	"net/http"
)

var (
	Message string
)

type Response struct {
	Queue   []pvisor.Packet
	Success bool
}

var Queue []pvisor.Packet

func main() {

	Queue = []pvisor.Packet{}
	router := mux.NewRouter()
	router.HandleFunc("/read/", read)
	router.HandleFunc("/write/", write)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./httpdocs"))))

	http.Handle("/", router)

	http.ListenAndServe("0.0.0.0:80", nil)
}

func write(w http.ResponseWriter, request *http.Request) {

	body, _ := ioutil.ReadAll(request.Body)
	var packet pvisor.Packet
	json.Unmarshal(body, &packet)

	Queue = append(Queue, packet)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{success:true}"))
}

func read(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if len(Queue) > 0 {
		resp := Response{Success: true, Queue: Queue}
		b, err := json.Marshal(resp)
		if err == nil {
			w.Write(b)
			Queue = []pvisor.Packet{}
		}

	} else {
		w.Write([]byte("{success:false}"))
	}
}
