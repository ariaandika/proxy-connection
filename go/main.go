package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
  signleHostProxy()
}

func signleHostProxy() {
  url, urlErr := url.Parse("http://localhost:8000")
  httputil.NewSingleHostReverseProxy(url)

  if urlErr != nil {
    log.Fatal("Invalid url", url)
  }

  httpErr := http.ListenAndServe(":3000", nil)
  
  if httpErr != nil {
    log.Fatal("Failed to listen and serve ", httpErr)
  }
}

