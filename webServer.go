package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    fmt.Printf("\nServer Running on port 8080!")
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/api/posts", Posts)
    router.HandleFunc("/api/posts/{postId}", ShowPost)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Access Denied")
}

func Posts(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home")
}

func ShowPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postId := vars["postId"]
    fmt.Fprintf(w, "Show post", postId)
}