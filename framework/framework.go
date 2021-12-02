package framework

import (
	"encoding/json"
	"hexagonal/app"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u app.User
	var a app.App
	if err := dec.Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.RegisterUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServeHTTP() {
	http.HandleFunc("/register_user", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
