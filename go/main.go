package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {

  port := getPort()

  signleHostProxy(port)
}

func signleHostProxy(port string) {
  url, urlErr := url.Parse("http://localhost:8000")

  if urlErr != nil {
    log.Fatal("Invalid url", url)
  }

  url.Host = url.Hostname() + ":" + port

  proxy := httputil.NewSingleHostReverseProxy(url)

  httpErr := http.ListenAndServe(":3000", proxy)
  
  if httpErr != nil {
    log.Fatal("Failed to listen and serve ", httpErr)
  }
}

func getPort() string {
  envs := os.Environ()

  for _, elem := range envs {
    s := strings.Split(elem, "=")
    if s[0] == "TARGET" {
      return s[1]
    }
  }

  log.Fatal("must provide TARGET env vars")
  return ""

  // file, err := os.ReadFile("../config")
  // if err != nil { log.Fatal(err) }
  //
  // content := string(file)
  //
  // port := strings.Split(content, "\n")[0]
  //
  // return port
}
