package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users []User

func main() {
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            var u User
            json.NewDecoder(r.Body).Decode(&u)
            u.ID = len(users) + 1
            users = append(users, u)
            json.NewEncoder(w).Encode(u)
        } else {
            json.NewEncoder(w).Encode(users)
        }
    })
    http.ListenAndServe(":8080", nil)
}
