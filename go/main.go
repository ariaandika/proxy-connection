package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
  singleHostProxy()
}

func singleHostProxy() {
  cert, key := getTls()
  url := getProxyUrl()
  proxy := httputil.NewSingleHostReverseProxy(url)
  
  if cert != "" && key != "" {
    log.Fatal(http.ListenAndServeTLS(getPort(), cert, key, proxy))
  } else {
    log.Fatal(http.ListenAndServe(getPort(), proxy))
  }
}

func panic(val *url.URL, err error) *url.URL {
  if err != nil {
    log.Fatal("Invalid url", val)
  }

  return val
}

func getProxyUrl() *url.URL {
  envs := os.Environ()

  u := ""

  for _, elem := range envs {
    s := strings.Split(elem, "=")
    if s[0] == "TARGET" {
      u = fmt.Sprintf("http://localhost:%s", s[1])
    }
  }

  if u == "" {
    log.Fatal("must provide TARGET env vars")
  }

  ur, err := url.Parse(u)

  if err != nil {
    log.Fatal("Invalid url", ur)
  }

  return ur
}

func getPort() string {
  envs := os.Environ()

  for _, elem := range envs {
    s := strings.Split(elem, "=")
    if s[0] == "PORT" {
      return fmt.Sprintf(":%s", s[1])
    }
  }

  return ":3000"
}

func getTls() (string, string) {
  envs := os.Environ()

  cert := ""
  key := ""

  for _, elem := range envs {
    s := strings.Split(elem, "=")
    if s[0] == "CERT" {
      cert = s[1]
    } else if s[0] == "KEY" {
      key = s[1]
    }
  }

  return cert, key
}
