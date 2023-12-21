# Proxy Connection

- if the backend server only support HTTP/1.1, should the proxy also use  
  HTTP/1.1 through the client ?

in nginx, thats the case

- should we even use HTTP/2.0 for server to server inside private network ?

actix only enable HTTP/2.0 when using ssl,
do we need ssl inside a private network ?

## Subject

bun, backend with only HTTP/1.1
rust, backend that support HTTP/2.0 via actix
go, the easiest reverse proxy server and tls

