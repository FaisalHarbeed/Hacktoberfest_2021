package main

import (
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("source"))
    http.Handle("/", fs) 

    http.ListenAndServe(":9527", nil)
}
