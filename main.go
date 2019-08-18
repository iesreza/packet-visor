package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"packet-visor/helper"
)

var (
	Message  string
)

type Response struct{
	 Queue []helper.Packet
	 Success bool
}

var Queue []helper.Packet

func main()  {

	Queue = []helper.Packet{}
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./httpdocs"))))
	http.Handle("/", router)

	router.HandleFunc("/read",read)

	http.ListenAndServe("0.0.0.0:80", nil)
}


func read(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if len(Queue) > 0{
		resp := Response{ Success:true, Queue:Queue }
		b,err := json.Marshal(resp)
		if err == nil{
			w.Write(b)
		}

	}else{
		w.Write([]byte("{success:false}"))
	}
}

