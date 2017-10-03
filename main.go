package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	port = 8080
)

type ApiAiRequest struct {
	Result struct {
		Parameters struct {
			GivenName string `json:"given-name"`
		}
	}
}

func debug_request(w http.ResponseWriter, r *http.Request) {
	var req ApiAiRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(req.Result.Parameters.GivenName)

}

func main() {
	http.HandleFunc("/", debug_request)

	log.Println("Listening at port ", port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
