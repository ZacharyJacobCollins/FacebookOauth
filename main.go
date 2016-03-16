package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func handleFacebook(w http.ResponseWriter, r *http.Request) {
	//Code grabbed when logging into facebook
	code := r.URL.Query().Get("code")
	//Access token using code making this call
	accessTokReq := "conn string here!" + code
	client := &http.Client{}
	req, err := http.NewRequest("GET", accessTokReq, nil)
	if err != nil {
		panic(err)
	}
	resp, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//gets the byte array of the
	contents, err := ioutil.ReadAll(resp.Body)
	log.Println(string(contents))
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/facebook", handleFacebook)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":1337", nil)
}
