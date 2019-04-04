package main

import (
  "io"
  "net/http"
)

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "<body style='background-color: white'><h1>Hello Tekton!</h1></body>")
}

