package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting Service...")

	portEnv := os.Getenv("PORT")
	greetEnv := os.Getenv("GREETING")

	greeter := Greeter{
		greeting: greetEnv,
	}

	http.HandleFunc("/", greeter.GreetingHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", portEnv), nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

type Greeter struct {
	greeting string
}

type GreetingResponse struct {
	Message string `json:"message"`
}

func (g *Greeter) GreetingHandler(w http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	name := params.Get("name")

	resp := GreetingResponse{
		Message: fmt.Sprintf("%s %s! Have an ok night!", g.greeting, name),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)

	print, e := json.MarshalIndent(resp, "", "\t")
	if e != nil {
		ServerError(w, e)
		return
	}

	log.Println(string(print))
}

func ServerError(w http.ResponseWriter, err error) {
	log.Println(err)

	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
