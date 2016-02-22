package main

import (
	"bytes"
	"log"
	"net/http"
)

func sendCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	client := &http.Client{}
	var jsonStr = []byte("grant_type=authorization_code&redirect_uri=http://zacc.xyz:8000&code=" + code)
	req, err := http.NewRequest("POST", "https://hackillinois.climate.com/api/oauth/token", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Basic ZHBxazVzbXBxMDM5Mmo6dDB0czB0YWdvcm05bnExdjZzbW10dnBxYzI=")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	res, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res)
	log.Println(res.Body)
}

func dashHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func handleFacebook(w http.ResponseWriter, r *http.Request) {
	log.Print("420")
	code := r.URL.Query().Get("code")
	log.Print(code)
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/facebook", handleFacebook)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":1337", nil)
}
