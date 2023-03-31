package main

import (
  "fmt"
  "net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is some code")
}

func main() {
  http.HandleFunc("/", greet)
  http.ListenAndServe(":8080", nil)
}
