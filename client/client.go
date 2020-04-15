package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var url string

func callApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("api.gtpl")
		fmt.Println("log", err)
		t.Execute(w, nil)
	} else {
		result := CallServerApi()
		fmt.Fprintf(w, result)

	}
}

func main() {
	url = os.Getenv("SERVER_URL")
	log.Println("server url", url)
	http.HandleFunc("/", callApi)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func CallServerApi() string {
	// url := fmt.Sprintf(url)
	log.Println("server url", url)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating http request", err)
	}
	request.Header.Set("content-type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("Error", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading resp body:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Err: request not proceeded", string(body))
	}
	log.Print("Result from server", string(body))
	return string(body)
}
