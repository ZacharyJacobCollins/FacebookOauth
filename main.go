package main

import (
	"log"
	"net/http"
)

func handleFacebook(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://graph.facebook.com/v2.3/oauth/access_token?client_id=1698453937058092&redirect_uri=http://zacc.xyz/facebook&client_secret=77d2a0169b92e291bd1a61837953973b&code="+code, nil)
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
