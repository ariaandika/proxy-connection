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
  port, cert, key, hosts := setupConfig()

  // proxy := httputil.NewSingleHostReverseProxy(targetUrl)

  proxy := &httputil.ReverseProxy{
    Rewrite: func(r *httputil.ProxyRequest) {
      if hosts[r.In.Host] != nil {
        r.SetURL(hosts[r.In.Host])
      }
    },
  }

  fmt.Printf("[go    public] Serving %s\n", port)
  
  if cert != "" && key != "" {
    fmt.Println("tls enabled")
    log.Fatal(http.ListenAndServeTLS(port, cert, key, proxy))
  } else {
    log.Fatal(http.ListenAndServe(port, proxy))
  }
}


func setupConfig() (string, string, string, map[string]*url.URL) {
  envs := os.Environ()
  
  port := ":3000"
  cert := ""
  key := ""

  hosts := make(map[string]*url.URL)

  for _, elem := range envs {
    s := strings.Split(elem, "=")

    if s[0] == "CERT" {
      cert = s[1]
    } else if s[0] == "KEY" {
      key = s[1]
    } else if s[0] == "PORT" {
      port = fmt.Sprint(":", s[1])
    // } else if s[0] == "TARGET" {
    //   u = fmt.Sprintf("http://localhost:%s", s[1])
    } else if strings.HasPrefix(s[0], "TARGET") {
      host := strings.Split(s[1], ":")

      ur, err := url.Parse(fmt.Sprintf("http://127.0.0.1:%s", host[1]))

      if err != nil {
        log.Fatal("Invalid target url", err)
      }

      hosts[host[0]] = ur
    }
  }

  return port, cert, key, hosts
}

