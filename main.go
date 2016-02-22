package main

import (
	"encoding/base64"
	"log"
	"net/http"
)

func handleFacebook(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	url := []byte("https://graph.facebook.com/v2.3/oauth/access_token?client_id=1698453937058092&redirect_uri=http://zacc.xyz:1337/facebook&client_secret=77d2a0169b92e291bd1a61837953973b&code=" + code)
	str := base64.StdEncoding.EncodeToString(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", str, nil)
	if err != nil {
		panic(err)
	}
	resp, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Body)
	log.Println(resp)
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/facebook", handleFacebook)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":1337", nil)
}
