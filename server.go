package main

import (
 "encoding/json";
 "fmt"
 "io/ioutil"
 "log"
 "net/http"
 "os"
)

type InvokeRequest struct {
  Data     map[string]json.RawMessage
  Metadata map[string]interface{}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write([]byte("hello world"))
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(body)
	}
}

func queueHandler(w http.ResponseWriter, r *http.Request) {
  var invokeRequest InvokeRequest

  d := json.NewDecoder(r.Body)
  d.Decode(&invokeRequest)

  var parsedMessage string
  json.Unmarshal(invokeRequest.Data["queueItem"], &parsedMessage)
  fmt.Println(parsedMessage) // your message
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)
    mux.HandleFunc("/queueTrigger", queueHandler)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
